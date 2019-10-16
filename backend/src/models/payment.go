package models

// PaymentStore interface type
type PaymentStore interface {
	Save(*Payment) error
	Delete(*Payment) error
}

// Payment objects
type Payment struct {
	store     PaymentStore
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	Amount    int 	  `json:"amount"`
	Timestamp int     `json:"timestamp"`
}

// NewPayment creates a new Payment object
func NewPayment(cs PaymentStore) *Payment {
	return &Payment{
		store: cs,
	}
}

// IsNew check if Payment object is new
func (c *Payment) IsNew() bool {
	return c.ID == 0
}

// Save is the function to save a Payment object
func (c *Payment) Save() error {
	return c.store.Save(c)
}

// Delete is the function to delete a Payment object
func (c *Payment) Delete() error {
	return c.store.Delete(c)
}
