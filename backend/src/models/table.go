package models

// TableStore interface type
type TableStore interface {
	Save(*Table) error
	Delete(*Table) error
}

// Table objects
type Table struct {
	store TableStore
	ID    int `json:"id"`
	Num   int `json:"table_num"`
}

// NewTable creates a new Table object
func NewTable(cs TableStore) *Table {
	return &Table{
		store: cs,
	}
}

// IsNew check if Table object is new
func (c *Table) IsNew() bool {
	return c.ID == 0
}

// Save is the function to save a Table object
func (c *Table) Save() error {
	return c.store.Save(c)
}

// Delete is the function to delete a Table object
func (c *Table) Delete() error {
	return c.store.Delete(c)
}
