package stores

import (
	"encoding/json"
	"errors"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

// OrderItemStore interface type for orderItemStore
type OrderItemStore interface {
	Get() ([]*models.OrderItem, error)
	GetByID(id int) (*models.OrderItem, error)
	GetByOrder(id int) ([]*models.OrderItem, error)
	Save(c *models.OrderItem) error
	Delete(c *models.OrderItem) error
}

// OrderItem struct type for orderItemStore
type OrderItem struct {
	e *common.Environment
}

// NewOrderItemStore function for orderItemStore
func NewOrderItemStore(e *common.Environment) *OrderItem {
	return &OrderItem{
		e: e,
	}
}

// Get function for orderItemStore
func (s *OrderItem) Get() ([]*models.OrderItem, error) {
	return s.getItemsFromDatabase("")
}

// GetByID function for orderItemStore
func (s *OrderItem) GetByID(id int) (*models.OrderItem, error) {
	if id == 0 {
		return nil, errors.New("Order item ID required")
	}

	sql := `WHERE "id" = ?`
	categories, err := s.getItemsFromDatabase(sql, id)
	if len(categories) == 0 {
		return nil, err
	}
	return categories[0], err
}

// GetByOrder function for orderItemStore
func (s *OrderItem) GetByOrder(id int) ([]*models.OrderItem, error) {
	if id == 0 {
		return nil, errors.New("Order ID required")
	}

	sql := `WHERE "order_id" = ?`
	return s.getItemsFromDatabase(sql, id)
}

func (s *OrderItem) getItemsFromDatabase(where string, values ...interface{}) ([]*models.OrderItem, error) {
	sql := `SELECT
				"id",
				"products",
				"order_id"
			FROM "order_item" ` + where

	rows, err := s.e.DB.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.OrderItem
	for rows.Next() {
		var id int
		var productsJSON string
		var orderID int

		err := rows.Scan(
			&id,
			&productsJSON,
			&orderID,
		)
		if err != nil {
			continue
		}

		var products []int
		json.Unmarshal([]byte(productsJSON), &products)

		orderItem := models.NewOrderItem(s)
		orderItem.ID = id
		orderItem.Products = products
		orderItem.OrderID = orderID

		results = append(results, orderItem)
	}
	return results, nil
}

// Save function for orderItemStore
func (s *OrderItem) Save(c *models.OrderItem) error {
	if c.ID == 0 {
		return s.saveNew(c)
	}
	return s.updateExisting(c)
}

func (s *OrderItem) updateExisting(c *models.OrderItem) error {
	sql := `UPDATE "order_item"
			SET "products" = ?,
				"order_id" = ?
			WHERE "id" = ?`

	productJSON, _ := json.Marshal(c.Products)

	_, err := s.e.DB.Exec(
		sql,
		productJSON,
		c.OrderID,
		c.ID,
	)
	return err
}

func (s *OrderItem) saveNew(c *models.OrderItem) error {
	sql := `INSERT INTO "order_item" ("products", "order_id") VALUES (?,?)`

	productJSON, _ := json.Marshal(c.Products)

	result, err := s.e.DB.Exec(sql, productJSON, c.OrderID)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	c.ID = int(id)
	return nil
}

// Delete function for orderItemStore
func (s *OrderItem) Delete(c *models.OrderItem) error {
	if c.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "order_item" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, c.ID)
	return err
}
