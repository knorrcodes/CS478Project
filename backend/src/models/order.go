package models

import (
	"time"
)

// OrderStore interface type
type OrderStore interface {
	Save(*Order) error
	Delete(*Order) error
}

// Order objects
type Order struct {
	store     OrderStore
	ID        int       `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	TableID   int       `json:"table_id"`
	ServerID  int       `json:"server_id"`
}

// NewOrder creates a new Order object
func NewOrder(cs OrderStore) *Order {
	return &Order{
		store: cs,
	}
}

// IsNew check if Order object is new
func (c *Order) IsNew() bool {
	return c.ID == 0
}

// IsOpen checks if Order object is currently open
func (c *Order) IsOpen() bool {
	return c.EndTime.Unix() == 0
}

// Save is the function to save a Order object
func (c *Order) Save() error {
	return c.store.Save(c)
}

// Delete is the function to delete a Order object
func (c *Order) Delete() error {
	return c.store.Delete(c)
}
