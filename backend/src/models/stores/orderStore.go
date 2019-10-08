package stores

import (
	"errors"
	"time"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

type OrderStatus int

const (
	OrderStatusOpened OrderStatus = iota
	OrderStatusClosed
	OrderStatusAny
)

type OrderStore interface {
	GetOrders(status OrderStatus) ([]*models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	GetOrdersByServer(id int, status OrderStatus) ([]*models.Order, error)
	GetOrdersByTable(id int, status OrderStatus) ([]*models.Order, error)
	GetLatestOrderByTable(id int) (*models.Order, error)
	Save(c *models.Order) error
	Delete(c *models.Order) error
}

type Order struct {
	e *common.Environment
}

func NewOrderStore(e *common.Environment) *Order {
	return &Order{
		e: e,
	}
}

func (s *Order) GetOrders(status OrderStatus) ([]*models.Order, error) {
	if status < OrderStatusOpened || status > OrderStatusAny {
		return nil, errors.New("invalid order status")
	}

	switch status {
	case OrderStatusClosed:
		sql := `WHERE "endtime" != 0`
		return s.getOrdersFromDatabase(sql)
	case OrderStatusOpened:
		sql := `WHERE "endtime" = 0`
		return s.getOrdersFromDatabase(sql)
	default:
		return s.getOrdersFromDatabase("")
	}
}

func (s *Order) GetOrderByID(id int) (*models.Order, error) {
	if id == 0 {
		return nil, errors.New("Order ID required")
	}

	sql := `WHERE "id" = ?`
	codes, err := s.getOrdersFromDatabase(sql, id)
	if len(codes) == 0 {
		return nil, err
	}
	return codes[0], err
}

func (s *Order) GetOrdersByServer(id int, status OrderStatus) ([]*models.Order, error) {
	if id == 0 {
		return nil, errors.New("Server ID required")
	}

	switch status {
	case OrderStatusClosed:
		sql := `WHERE "server_id" = ? AND "endtime" != 0`
		return s.getOrdersFromDatabase(sql, id)
	case OrderStatusOpened:
		sql := `WHERE "server_id" = ? AND "endtime" = 0`
		return s.getOrdersFromDatabase(sql, id)
	default:
		sql := `WHERE "server_id" = ?`
		return s.getOrdersFromDatabase(sql, id)
	}
}

func (s *Order) GetOrdersByTable(id int, status OrderStatus) ([]*models.Order, error) {
	if id == 0 {
		return nil, errors.New("Table ID required")
	}

	switch status {
	case OrderStatusClosed:
		sql := `WHERE "table_id" = ? AND "endtime" != 0`
		return s.getOrdersFromDatabase(sql, id)
	case OrderStatusOpened:
		sql := `WHERE "table_id" = ? AND "endtime" = 0`
		return s.getOrdersFromDatabase(sql, id)
	default:
		sql := `WHERE "table_id" = ?`
		return s.getOrdersFromDatabase(sql, id)
	}
}

func (s *Order) GetLatestOrderByTable(id int) (*models.Order, error) {
	if id == 0 {
		return nil, errors.New("Table ID required")
	}

	sql := `WHERE "table_id" = ? ORDER BY "starttime" DESC LIMIT 1`
	orders, err := s.getOrdersFromDatabase(sql, id)
	if len(orders) == 0 {
		return nil, err
	}
	return orders[0], err
}

func (s *Order) getOrdersFromDatabase(where string, values ...interface{}) ([]*models.Order, error) {
	sql := `SELECT
				"id",
				"starttime",
				"endtime",
				"table_id",
				"server_id"
			FROM "order" ` + where

	rows, err := s.e.DB.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Order
	for rows.Next() {
		var id int
		var startTime int64
		var endTime int64
		var tableID int
		var serverID int

		err := rows.Scan(
			&id,
			&startTime,
			&endTime,
			&tableID,
			&serverID,
		)
		if err != nil {
			continue
		}

		Order := models.NewOrder(s)
		Order.ID = id
		Order.StartTime = time.Unix(startTime, 0)
		Order.EndTime = time.Unix(endTime, 0)
		Order.TableID = tableID
		Order.ServerID = serverID

		results = append(results, Order)
	}
	return results, nil
}

func (s *Order) Save(c *models.Order) error {
	if c.ID == 0 {
		return s.saveNew(c)
	}
	return s.updateExisting(c)
}

func (s *Order) updateExisting(c *models.Order) error {
	sql := `UPDATE "order"
			SET "starttime" = ?,
				"endtime" = ?,
				"table_id" = ?,
				"server_id" = ?
			WHERE "id" = ?`

	_, err := s.e.DB.Exec(
		sql,
		c.StartTime.Unix(),
		c.EndTime.Unix(),
		c.TableID,
		c.ServerID,
		c.ID,
	)
	return err
}

func (s *Order) saveNew(c *models.Order) error {
	sql := `INSERT INTO "order" (
				"starttime",
				"endtime",
				"table_id",
				"server_id"
			) VALUES (?, ?, ?, ?)`

	result, err := s.e.DB.Exec(
		sql,
		c.StartTime.Unix(),
		c.EndTime.Unix(),
		c.TableID,
		c.ServerID,
	)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	c.ID = int(id)
	return nil
}

func (s *Order) Delete(c *models.Order) error {
	if c.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "order" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, c.ID)
	return err
}
