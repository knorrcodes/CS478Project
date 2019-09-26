package stores

import (
	"errors"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

type CategoryStore interface {
	GetCategories() ([]*models.Category, error)
	GetCategoryByID(id int) (*models.Category, error)
	Save(c *models.Category) error
	Delete(c *models.Category) error
}

type Category struct {
	e *common.Environment
}

func NewCategoryStore(e *common.Environment) *Category {
	return &Category{
		e: e,
	}
}

func (s *Category) GetCategories() ([]*models.Category, error) {
	return s.getCategoriesFromDatabase("")
}

func (s *Category) GetCategoryByID(id int) (*models.Category, error) {
	if id == 0 {
		return nil, errors.New("Category ID required")
	}

	sql := `WHERE "id" = ?`
	categories, err := s.getCategoriesFromDatabase(sql, id)
	if len(categories) == 0 {
		return nil, err
	}
	return categories[0], err
}

func (s *Category) getCategoriesFromDatabase(where string, values ...interface{}) ([]*models.Category, error) {
	sql := `SELECT
				"id",
				"name"
			FROM "category" ` + where

	rows, err := s.e.DB.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Category
	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(
			&id,
			&name,
		)
		if err != nil {
			continue
		}

		category := models.NewCategory(s)
		category.ID = id
		category.Name = name

		results = append(results, category)
	}
	return results, nil
}

func (s *Category) Save(c *models.Category) error {
	if c.ID == 0 {
		return s.saveNew(c)
	}
	return s.updateExisting(c)
}

func (s *Category) updateExisting(c *models.Category) error {
	sql := `UPDATE "category"
			SET "name" = ?
			WHERE "id" = ?`

	_, err := s.e.DB.Exec(
		sql,
		c.Name,
		c.ID,
	)
	return err
}

func (s *Category) saveNew(c *models.Category) error {
	if c.Name == "" {
		return errors.New("Category name cannot be empty")
	}

	sql := `INSERT INTO "category" ("name") VALUES (?)`

	result, err := s.e.DB.Exec(sql, c.Name)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	c.ID = int(id)
	return nil
}

func (s *Category) Delete(c *models.Category) error {
	if c.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "category" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, c.ID)
	return err
}
