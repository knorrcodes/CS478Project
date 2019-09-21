package stores

import (
	"errors"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

type ProductStore interface {
	GetProductByID(id int) (*models.Product, error)
	Save(u *models.Product) error
	Delete(u *models.Product) error
}

type Product struct {
	e *common.Environment
}

func NewProductStore(e *common.Environment) *Product {
	return &Product{
		e: e,
	}
}

func (s *Product) GetProductByID(id int) (*models.Product, error) {
	if id == 0 {
		return nil, errors.New("Product ID required")
	}

	sql := `WHERE "id" = ?`
	users, err := s.getProductsFromDatabase(sql, id)
	if len(users) == 0 {
		return nil, err
	}
	return users[0], err
}

func (s *Product) getProductsFromDatabase(where string, values ...interface{}) ([]*models.Product, error) {
	sql := `SELECT
				"id",
				"name",
				"desc",
				"picture",
				"price",
				"category",
				"ws_cost",
				"num_of_sides"
			FROM "product" ` + where

	rows, err := s.e.DB.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Product
	for rows.Next() {
		var id int
		var name string
		var desc string
		var picture string
		var price int
		var category int
		var wsCost int
		var numOfSides int

		err := rows.Scan(
			&id,
			&name,
			&desc,
			&picture,
			&price,
			&category,
			&wsCost,
			&numOfSides,
		)
		if err != nil {
			continue
		}

		product := models.NewProduct(s.e, s)
		product.ID = id
		product.Name = name
		product.Desc = desc
		product.Picture = picture
		product.Price = price
		product.Category = category
		product.WSCost = wsCost
		product.NumOfSides = numOfSides

		results = append(results, product)
	}
	return results, nil
}

func (s *Product) Save(p *models.Product) error {
	if p.ID == 0 {
		return s.saveNew(p)
	}
	return s.updateExisting(p)
}

func (s *Product) updateExisting(p *models.Product) error {
	return nil
}

func (s *Product) saveNew(p *models.Product) error {
	return nil
}

func (s *Product) Delete(p *models.Product) error {
	if p.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "product" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, p.ID)
	return err
}
