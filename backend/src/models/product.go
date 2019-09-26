package models

type ProductStore interface {
	Save(*Product) error
	Delete(*Product) error
}

type Product struct {
	store      ProductStore
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Picture    string `json:"picture"`
	Price      int    `json:"price"`
	Category   int    `json:"category"`
	WSCost     int    `json:"wscost"`
	NumOfSides int    `json:"num_of_sides"`
}

// NewProduct creates a new product object
func NewProduct(ps ProductStore) *Product {
	return &Product{
		store: ps,
	}
}

func (p *Product) IsNew() bool {
	return p.ID == 0
}

func (p *Product) Save() error {
	return p.store.Save(p)
}

func (p *Product) Delete() error {
	return p.store.Delete(p)
}
