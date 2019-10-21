package models

// OrderItemStore interface type
type OrderItemStore interface {
	Save(*OrderItem) error
	Delete(*OrderItem) error
}

// OrderItem objects
type OrderItem struct {
	store    OrderItemStore
	ID       int   `json:"id"`
	OrderID  int   `json:"order_id"`
	Products []int `json:"products"`
}

// NewOrderItem creates a new OrderItem object
func NewOrderItem(cs OrderItemStore) *OrderItem {
	return &OrderItem{
		store: cs,
	}
}

// IsNew check if OrderItem object is new
func (c *OrderItem) IsNew() bool {
	return c.ID == 0
}

// Save is the function to save a OrderItem object
func (c *OrderItem) Save() error {
	return c.store.Save(c)
}

// Delete is the function to delete a OrderItem object
func (c *OrderItem) Delete() error {
	return c.store.Delete(c)
}
