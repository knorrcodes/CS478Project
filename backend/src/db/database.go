package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql" // MySQL driver
	log "github.com/lfkeitel/verbose/v5"

	"koala.pos/src/common"
)

type database struct {
	createFuncs  map[string]func(*common.DatabaseAccessor) error
	migrateFuncs []migrateFunc
}

func newDBConnector() *database {
	m := &database{}

	m.createFuncs = map[string]func(*common.DatabaseAccessor) error{
		"product":    m.createProductTable,
		"category":   m.createCategoryTable,
		"server":     m.createServerTable,
		"order":      m.createOrderTable,
		"table":      m.createTableTable,
		"cust_code":  m.createCustCodeTable,
		"order_item": m.createOrderItemTable,
		"settings":   m.createSettingTable,
	}

	m.migrateFuncs = []migrateFunc{}

	return m
}

func (m *database) init(d *common.DatabaseAccessor, c *common.Config) error {
	if err := m.connect(d, c); err != nil {
		return err
	}
	d.Driver = "mysql"

	if err := m.createTables(d); err != nil {
		return err
	}

	return m.migrateTables(d, c)
}

func (m *database) connect(d *common.DatabaseAccessor, c *common.Config) error {
	if c.Database.Port == 0 {
		c.Database.Port = 3306
	}

	mc := mysql.NewConfig()
	mc.User = c.Database.Username
	mc.Passwd = c.Database.Password
	mc.Net = "tcp"
	mc.Addr = fmt.Sprintf("%s:%d", c.Database.Address, c.Database.Port)
	mc.DBName = c.Database.Name

	var err error
	d.DB, err = sql.Open("mysql", mc.FormatDSN())
	if err != nil {
		return err
	}

	if err := d.DB.Ping(); err != nil {
		return err
	}

	// Check the SQL mode, the user is responsible for setting it
	row := d.DB.QueryRow(`SELECT @@GLOBAL.sql_mode`)

	mode := ""
	if err := row.Scan(&mode); err != nil {
		return err
	}

	if !strings.Contains(mode, "ANSI") {
		return errors.New("MySQL must be in ANSI mode. Please set the global mode or edit the my.cnf file to enable ANSI sql_mode.")
	}
	return nil
}

func (m *database) createTables(d *common.DatabaseAccessor) error {
	rows, err := d.DB.Query(`SHOW TABLES`)
	if err != nil {
		return err
	}
	defer rows.Close()

	tables := make(map[string]bool)
	for _, table := range common.DatabaseTableNames {
		tables[table] = false
	}

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}
		tables[tableName] = true
	}

	for table, create := range m.createFuncs {
		if !tables[table] {
			fmt.Printf("Creating table %s\n", table)
			if err := create(d); err != nil {
				return err
			}
		}
	}

	return m.buildConstraints(d)
}

func (m *database) migrateTables(d *common.DatabaseAccessor, c *common.Config) error {
	var currDBVer int
	verRow := d.DB.QueryRow(`SELECT "value" FROM "settings" WHERE "id" = 'db_version'`)
	if verRow == nil {
		return errors.New("Failed to get database version")
	}
	verRow.Scan(&currDBVer)

	log.WithFields(log.Fields{
		"current-version": currDBVer,
		"active-version":  DBVersion,
	}).Debug("Database Versions")

	// No migration needed
	if currDBVer == DBVersion {
		return nil
	}

	if currDBVer > DBVersion {
		return errors.New("Database is too new, can't rollback")
	}

	neededMigrations := m.migrateFuncs[currDBVer:DBVersion]
	for _, migrate := range neededMigrations {
		if migrate == nil {
			continue
		}
		if err := migrate(d, c); err != nil {
			return err
		}
	}

	_, err := d.DB.Exec(`UPDATE "settings" SET "value" = ? WHERE "id" = 'db_version'`, DBVersion)
	return err
}

func (m *database) createSettingTable(d *common.DatabaseAccessor) error {
	sql := `CREATE TABLE "settings" (
		"id" VARCHAR(255) PRIMARY KEY NOT NULL,
		"value" TEXT NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	if _, err := d.DB.Exec(sql); err != nil {
		return err
	}

	_, err := d.DB.Exec(`INSERT INTO "settings" ("id", "value") VALUES ('db_version', ?)`, DBVersion)
	return err
}

func (m *database) createProductTable(d *common.DatabaseAccessor) error {
	sql := `CREATE TABLE "product" (
		"id" INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		"name" TINYTEXT NOT NULL,
		"desc" TEXT NOT NULL DEFAULT (''),
		"picture" TINYTEXT NOT NULL DEFAULT (''),
		"price" INT NOT NULL,
		"category_id" INT NOT NULL,
		"ws_cost" INT NOT NULL,
		"num_of_sides" TINYINT NOT NULL DEFAULT (0),
		INDEX ("category_id")
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1`

	if _, err := d.DB.Exec(sql); err != nil {
		return err
	}

	_, err := d.DB.Exec(`INSERT INTO "product"
		("name", "price", "category_id", "ws_cost", "num_of_sides")
		VALUES
		('Teriyaki Chicken', 1250, 1, 0, 2),
		('Grilled Chicken', 1150, 1, 0, 2),
		('Chicken Salad', 899, 1, 0, 1),
		('Chicken Tenders', 1099, 1, 0, 1),

		('6oz Filet', 1199, 2, 0, 2),
		('9oz Filet', 1499, 2, 0, 2),
		('Steak Salad', 1199, 2, 0, 1),

		('Wild Salmon', 1399, 3, 0, 2),
		('Catfish Dinner', 1199, 3, 0, 2),
		('Grilled Shrimp', 1299, 3, 0, 2),
		('Lobster Tail', 1895, 3, 0, 2),

		('All American Cheese Burger', 955, 4, 0, 1),
		('Bacon Cheese Burger', 1155, 4, 0, 1),
		('Veggie Burger', 1095, 4, 0, 1),

		('Fried Pickles', 815, 5, 0, 0),
		('Mozzarella Sticks', 735, 5, 0, 0),

		('Epic Brownie', 650, 6, 0, 0),
		('Fried Ice Cream', 550, 6, 0, 0),

		('French Fries', 150, 7, 0, 0),
		('Baked Potato', 150, 7, 0, 0),
		('Broccoli', 150, 7, 0, 0),
		('Rice', 150, 7, 0, 0),
		('Asparagus', 150, 7, 0, 0),

		('Coke', 200, 8, 0, 0),
		('Tea', 200, 8, 0, 0),
		('Water', 0, 8, 0, 0),
		('Lemonaid', 200, 8, 0, 0)`)
	return err
}

func (m *database) createCategoryTable(d *common.DatabaseAccessor) error {
	sql := `CREATE TABLE "category" (
		"id" INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		"name" TINYTEXT NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1`

	if _, err := d.DB.Exec(sql); err != nil {
		return err
	}

	_, err := d.DB.Exec(`INSERT INTO "category" ("id", "name") VALUES
							(1, 'Chicken'),
							(2, 'Steak'),
							(3, 'Fish'),
							(4, 'Burger'),
							(5, 'Appetizer'),
							(6, 'Dessert'),
							(7, 'Side'),
							(8, 'Drink')`)
	return err
}

func (m *database) createServerTable(d *common.DatabaseAccessor) error {
	sql := `CREATE TABLE "server" (
		"id" INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		"name" TINYTEXT NOT NULL,
		"code" MEDIUMINT NOT NULL,
		"manager" BOOLEAN NOT NULL DEFAULT(0)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1`

	if _, err := d.DB.Exec(sql); err != nil {
		return err
	}

	_, err := d.DB.Exec(`INSERT INTO "server" ("name", "code", "manager") VALUES
							('Default Server', 478, 0),
							('manager', 125, 1)`)
	return err
}

func (m *database) createOrderTable(d *common.DatabaseAccessor) error {
	sql := `CREATE TABLE "order" (
		"id" INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		"starttime" DATETIME NOT NULL,
		"endtime" DATETIME NOT NULL,
		"table_id" INT NOT NULL,
		"server_id" INT NOT NULL,
		INDEX ("table_id", "server_id")
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1`
	_, err := d.DB.Exec(sql)
	return err
}

func (m *database) createTableTable(d *common.DatabaseAccessor) error {
	sql := `CREATE TABLE "table" (
		"id" INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		"table_num" INT NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1`

	if _, err := d.DB.Exec(sql); err != nil {
		return err
	}

	_, err := d.DB.Exec(`INSERT INTO "table" ("table_num") VALUES
							(1),
							(2)`)
	return err
}

func (m *database) createCustCodeTable(d *common.DatabaseAccessor) error {
	sql := `CREATE TABLE "cust_code" (
		"id" INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		"starttime" DATETIME NOT NULL,
		"endtime" DATETIME NOT NULL,
		"code" TINYTEXT NOT NULL,
		"order_id" INT NOT NULL,
		INDEX ("order_id")
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1`
	_, err := d.DB.Exec(sql)
	return err
}

func (m *database) createOrderItemTable(d *common.DatabaseAccessor) error {
	sql := `CREATE TABLE "order_item" (
		"id" INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
		"products" JSON NOT NULL,
		"order_id" INT NOT NULL,
		INDEX ("order_id")
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1`
	_, err := d.DB.Exec(sql)
	return err
}

func (m *database) buildConstraints(d *common.DatabaseAccessor) error {
	alterStmts := []string{
		`ALTER TABLE "product"
			ADD FOREIGN KEY ("category_id")
			REFERENCES "category" ("id")`,
		`ALTER TABLE "order"
			ADD FOREIGN KEY ("table_id")
			REFERENCES "table" ("id")`,
		`ALTER TABLE "order"
			ADD FOREIGN KEY ("server_id")
			REFERENCES "server" ("id")`,
		`ALTER TABLE "cust_code"
			ADD FOREIGN KEY ("order_id")
			REFERENCES "order" ("id")`,
		`ALTER TABLE "order_item"
			ADD FOREIGN KEY ("order_id")
			REFERENCES "order" ("id")`,
	}

	for _, stmt := range alterStmts {
		if _, err := d.DB.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}
