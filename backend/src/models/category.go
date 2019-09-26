package models

type CategoryStore interface {
	Save(*Category) error
	Delete(*Category) error
}

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

func (c *Category) IsNew() bool {
	return c.ID == 0
}

func (c *Category) Save() error {
	return c.store.Save(c)
}

func (c *Category) Delete() error {
	return c.store.Delete(c)
}
