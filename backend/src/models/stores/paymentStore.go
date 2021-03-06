package stores

import (
	"errors"
	"time"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

// PaymentStore type for payment
type PaymentStore interface {
	GetPayments() ([]*models.Payment, error)
	GetPaymentByID(id int) (*models.Payment, error)
	GetPaymentsForOrder(orderID int) ([]*models.Payment, error)
	Save(c *models.Payment) error
	Delete(c *models.Payment) error
}

// Payment type for payment
type Payment struct {
	e *common.Environment
}

// NewPaymentStore function for payment
func NewPaymentStore(e *common.Environment) *Payment {
	return &Payment{
		e: e,
	}
}

// GetPayments function for payment
func (s *Payment) GetPayments() ([]*models.Payment, error) {
	return s.getPaymentsFromDatabase("")
}

// GetPaymentByID function for payment
func (s *Payment) GetPaymentByID(id int) (*models.Payment, error) {
	if id == 0 {
		return nil, errors.New("Payment ID required")
	}

	sql := `WHERE "id" = ? ORDER BY "timestamp"`
	payments, err := s.getPaymentsFromDatabase(sql, id)
	if len(payments) == 0 {
		return nil, err
	}
	return payments[0], err
}

//GetPaymentsForOrder function for payment
func (s *Payment) GetPaymentsForOrder(orderID int) ([]*models.Payment, error) {
	if orderID == 0 {
		return nil, errors.New("Order ID required")
	}

	sql := `WHERE "order_id" = ? ORDER BY "timestamp"`
	return s.getPaymentsFromDatabase(sql, orderID)
}

func (s *Payment) getPaymentsFromDatabase(where string, values ...interface{}) ([]*models.Payment, error) {
	sql := `SELECT
				"id",
				"order_id",
				"amount",
				"timestamp"
			FROM "payment" ` + where

	rows, err := s.e.DB.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Payment
	for rows.Next() {
		var id int
		var orderID int
		var amount int
		var timestamp int64

		err := rows.Scan(
			&id,
			&orderID,
			&amount,
			&timestamp,
		)
		if err != nil {
			continue
		}

		Payment := models.NewPayment(s)
		Payment.ID = id
		Payment.OrderID = orderID
		Payment.Amount = amount
		Payment.Timestamp = time.Unix(timestamp, 0)

		results = append(results, Payment)
	}
	return results, nil
}

// Save function for payment
func (s *Payment) Save(c *models.Payment) error {
	if c.ID == 0 {
		return s.saveNew(c)
	}
	return s.updateExisting(c)
}

func (s *Payment) updateExisting(c *models.Payment) error {
	sql := `UPDATE "payment"
			SET "order_id" = ?,
				"amount" = ?,
				"timestamp" = ?
			WHERE "id" = ?`

	_, err := s.e.DB.Exec(
		sql,
		c.OrderID,
		c.Amount,
		c.Timestamp.Unix(),
		c.ID,
	)
	return err
}

func (s *Payment) saveNew(c *models.Payment) error {
	if c.Amount == 0 {
		return errors.New("Payment Amount cannot be empty")
	}

	sql := `INSERT INTO "payment" (
				"order_id",
				"amount",
				"timestamp"
			) VALUES (?, ?, ?)`

	result, err := s.e.DB.Exec(
		sql,
		c.OrderID,
		c.Amount,
		c.Timestamp.Unix(),
	)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	c.ID = int(id)
	return nil
}

// Delete function for payment
func (s *Payment) Delete(c *models.Payment) error {
	if c.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "payment" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, c.ID)
	return err
}
