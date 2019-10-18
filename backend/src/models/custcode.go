package models

import "time"

// CustCodeStore interface type
type CustCodeStore interface {
	Save(*CustCode) error
	Delete(*CustCode) error
}

// CustCode objects
type CustCode struct {
	store     CustCodeStore
	ID        int       `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Code      string    `json:"code"`
	OrderID   int       `json:"order_id"`
}

// NewCustCode creates a new CustCode object
func NewCustCode(cs CustCodeStore) *CustCode {
	return &CustCode{
		store: cs,
	}
}

// IsNew check if CustCode object is new
func (c *CustCode) IsNew() bool {
	return c.ID == 0
}

// Save is the function to save a CustCode object
func (c *CustCode) Save() error {
	return c.store.Save(c)
}

// Delete is the function to delete a CustCode object
func (c *CustCode) Delete() error {
	return c.store.Delete(c)
}
