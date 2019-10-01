package stores

import (
	"errors"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

type TableStore interface {
	GetTables() ([]*models.Table, error)
	GetTableByID(id int) (*models.Table, error)
	GetTableByNumber(num int) (*models.Table, error)
	Save(c *models.Table) error
	Delete(c *models.Table) error
}

type Table struct {
	e *common.Environment
}

func NewTableStore(e *common.Environment) *Table {
	return &Table{
		e: e,
	}
}

func (s *Table) GetTables() ([]*models.Table, error) {
	return s.getTablesFromDatabase("")
}

func (s *Table) GetTableByID(id int) (*models.Table, error) {
	if id == 0 {
		return nil, errors.New("Table ID required")
	}

	sql := `WHERE "id" = ?`
	categories, err := s.getTablesFromDatabase(sql, id)
	if len(categories) == 0 {
		return nil, err
	}
	return categories[0], err
}

func (s *Table) GetTableByNumber(id int) (*models.Table, error) {
	if id == 0 {
		return nil, errors.New("Table number required")
	}

	sql := `WHERE "table_num" = ?`
	categories, err := s.getTablesFromDatabase(sql, id)
	if len(categories) == 0 {
		return nil, err
	}
	return categories[0], err
}

func (s *Table) getTablesFromDatabase(where string, values ...interface{}) ([]*models.Table, error) {
	sql := `SELECT
				"id",
				"table_num"
			FROM "table" ` + where

	rows, err := s.e.DB.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Table
	for rows.Next() {
		var id int
		var num int

		err := rows.Scan(
			&id,
			&num,
		)
		if err != nil {
			continue
		}

		Table := models.NewTable(s)
		Table.ID = id
		Table.Num = num

		results = append(results, Table)
	}
	return results, nil
}

func (s *Table) Save(c *models.Table) error {
	if c.ID == 0 {
		return s.saveNew(c)
	}
	return s.updateExisting(c)
}

func (s *Table) updateExisting(c *models.Table) error {
	sql := `UPDATE "table"
			SET "table_num" = ?
			WHERE "id" = ?`

	_, err := s.e.DB.Exec(
		sql,
		c.Num,
		c.ID,
	)
	return err
}

func (s *Table) saveNew(c *models.Table) error {
	if c.Num == 0 {
		return errors.New("Table number cannot be empty")
	}

	sql := `INSERT INTO "table" ("table_num") VALUES (?)`

	result, err := s.e.DB.Exec(sql, c.Num)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	c.ID = int(id)
	return nil
}

func (s *Table) Delete(c *models.Table) error {
	if c.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "table" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, c.ID)
	return err
}
