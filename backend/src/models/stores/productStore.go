package stores

import (
	"errors"
	"fmt"
	"strings"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

// ProductStore interface type for productStore
type ProductStore interface {
	GetProducts() ([]*models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	GetProductsByID(ids []int) ([]*models.Product, error)
	GetProductsByCategory(id int) ([]*models.Product, error)
	Save(p *models.Product) error
	Delete(p *models.Product) error
}

// Product struct type for productStore
type Product struct {
	e *common.Environment
}

// NewProductStore function for productStore
func NewProductStore(e *common.Environment) *Product {
	return &Product{
		e: e,
	}
}

// GetProducts function for productStore
func (s *Product) GetProducts() ([]*models.Product, error) {
	return s.getProductsFromDatabase("")
}

// GetProductByID function for productStore
func (s *Product) GetProductByID(id int) (*models.Product, error) {
	if id == 0 {
		return nil, errors.New("Product ID required")
	}

	sql := `WHERE "id" = ?`
	products, err := s.getProductsFromDatabase(sql, id)
	if len(products) == 0 {
		return nil, err
	}
	return products[0], err
}

// GetProductsByCategory function for productStore
func (s *Product) GetProductsByCategory(id int) ([]*models.Product, error) {
	if id <= 0 {
		return nil, errors.New("Category ID required")
	}

	sql := `WHERE "category_id" = ?`
	return s.getProductsFromDatabase(sql, id)
}

func intsToInterfaces(i []int) []interface{} {
	ret := make([]interface{}, len(i))
	for idx, v := range i {
		ret[idx] = v
	}
	return ret
}

// GetProductsByID function for productStore
func (s *Product) GetProductsByID(ids []int) ([]*models.Product, error) {
	if len(ids) == 0 {
		return nil, errors.New("Product IDs required")
	}

	whereIn := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")

	sql := fmt.Sprintf(`WHERE "id" IN (%s)`, whereIn)
	return s.getProductsFromDatabase(sql, intsToInterfaces(ids)...)
}

func (s *Product) getProductsFromDatabase(where string, values ...interface{}) ([]*models.Product, error) {
	sql := `SELECT
				"id",
				"name",
				"desc",
				"picture",
				"price",
				"category_id",
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

		product := models.NewProduct(s)
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

// Save function for productStore
func (s *Product) Save(p *models.Product) error {
	if p.ID == 0 {
		return s.saveNew(p)
	}
	return s.updateExisting(p)
}

func (s *Product) updateExisting(p *models.Product) error {
	sql := `UPDATE "product"
			SET	"name" = ?,
				"desc" = ?,
				"picture" = ?,
				"price" = ?,
				"category_id" = ?,
				"ws_cost" = ?,
				"num_of_sides" = ?
			WHERE "id" = ?`

	_, err := s.e.DB.Exec(
		sql,
		p.Name,
		p.Desc,
		p.Picture,
		p.Price,
		p.Category,
		p.WSCost,
		p.NumOfSides,
		p.ID,
	)
	return err
}

func (s *Product) saveNew(p *models.Product) error {
	if p.Name == "" {
		return errors.New("Product name cannot be empty")
	}

	sql := `INSERT INTO "product"
				("name",
				"desc",
				"picture",
				"price",
				"category_id",
				"ws_cost",
				"num_of_sides") VALUES (?,?,?,?,?,?,?)`

	result, err := s.e.DB.Exec(
		sql,
		p.Name,
		p.Desc,
		p.Picture,
		p.Price,
		p.Category,
		p.WSCost,
		p.NumOfSides,
	)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)
	return nil
}

// Delete function for productStore
func (s *Product) Delete(p *models.Product) error {
	if p.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "product" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, p.ID)
	return err
}
