package models

// CategoryStore interface type
type CategoryStore interface {
	Save(*Category) error
	Delete(*Category) error
}

// Category objects
type Category struct {
	store CategoryStore
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

// NewCategory creates a new Category object
func NewCategory(cs CategoryStore) *Category {
	return &Category{
		store: cs,
	}
}

// IsNew check if category object is new
func (c *Category) IsNew() bool {
	return c.ID == 0
}

// Save is the function to save a category object
func (c *Category) Save() error {
	return c.store.Save(c)
}

// Delete is the function to delete a category object
func (c *Category) Delete() error {
	return c.store.Delete(c)
}
