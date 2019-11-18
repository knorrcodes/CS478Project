package stores

import (
	"errors"
	"time"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

// CustCodeStore interface type for custCodeStore
type CustCodeStore interface {
	GetCustCodes() ([]*models.CustCode, error)
	GetCustCodeByID(id int) (*models.CustCode, error)
	GetCustCodeByCode(code string) (*models.CustCode, error)
	GetCustCodeByOrderID(id int) (*models.CustCode, error)
	Save(c *models.CustCode) error
	Delete(c *models.CustCode) error
}

// CustCode struct type for custCodeStore
type CustCode struct {
	e *common.Environment
}

// NewCustCodeStore function for custCodeStore
func NewCustCodeStore(e *common.Environment) *CustCode {
	return &CustCode{
		e: e,
	}
}

// GetCustCodes function for custCodeStore
func (s *CustCode) GetCustCodes() ([]*models.CustCode, error) {
	return s.getCodesFromDatabase("")
}

// GetCustCodeByID function for custCodeStore
func (s *CustCode) GetCustCodeByID(id int) (*models.CustCode, error) {
	if id == 0 {
		return nil, errors.New("CustCode ID required")
	}

	sql := `WHERE "id" = ?`
	codes, err := s.getCodesFromDatabase(sql, id)
	if len(codes) == 0 {
		return nil, err
	}
	return codes[0], err
}

// GetCustCodeByOrderID function for custCodeStore
func (s *CustCode) GetCustCodeByOrderID(id int) (*models.CustCode, error) {
	if id == 0 {
		return nil, errors.New("CustCode Order ID required")
	}

	sql := `WHERE "order_id" = ?`
	codes, err := s.getCodesFromDatabase(sql, id)
	if len(codes) == 0 {
		return nil, err
	}
	return codes[0], err
}

// GetCustCodeByCode function for custCodeStore
func (s *CustCode) GetCustCodeByCode(code string) (*models.CustCode, error) {
	if code == "" {
		return nil, nil
	}

	sql := `WHERE "code" = ?`
	codes, err := s.getCodesFromDatabase(sql, code)
	if len(codes) == 0 {
		return nil, err
	}
	return codes[0], err
}

func (s *CustCode) getCodesFromDatabase(where string, values ...interface{}) ([]*models.CustCode, error) {
	sql := `SELECT
				"id",
				"starttime",
				"endtime",
				"code",
				"order_id"
			FROM "cust_code" ` + where

	rows, err := s.e.DB.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.CustCode
	for rows.Next() {
		var id int
		var startTime int64
		var endTime int64
		var code string
		var orderID int

		err := rows.Scan(
			&id,
			&startTime,
			&endTime,
			&code,
			&orderID,
		)
		if err != nil {
			continue
		}

		CustCode := models.NewCustCode(s)
		CustCode.ID = id
		CustCode.StartTime = time.Unix(startTime, 0)
		CustCode.EndTime = time.Unix(endTime, 0)
		CustCode.Code = code
		CustCode.OrderID = orderID

		results = append(results, CustCode)
	}
	return results, nil
}

// Save function for custCodeStore
func (s *CustCode) Save(c *models.CustCode) error {
	if c.ID == 0 {
		return s.saveNew(c)
	}
	return s.updateExisting(c)
}

func (s *CustCode) updateExisting(c *models.CustCode) error {
	sql := `UPDATE "cust_code"
			SET "starttime" = ?,
				"endtime" = ?,
				"code" = ?,
				"order_id" = ?
			WHERE "id" = ?`

	_, err := s.e.DB.Exec(
		sql,
		c.StartTime.Unix(),
		c.EndTime.Unix(),
		c.Code,
		c.OrderID,
		c.ID,
	)
	return err
}

func (s *CustCode) saveNew(c *models.CustCode) error {
	if c.Code == "" {
		return errors.New("CustCode code cannot be empty")
	}

	sql := `INSERT INTO "cust_code" (
				"starttime",
				"endtime",
				"code",
				"order_id"
			) VALUES (?, ?, ?, ?)`

	result, err := s.e.DB.Exec(
		sql,
		c.StartTime.Unix(),
		c.EndTime.Unix(),
		c.Code,
		c.OrderID,
	)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	c.ID = int(id)
	return nil
}

// Delete function for custCodeStore
func (s *CustCode) Delete(c *models.CustCode) error {
	if c.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "cust_code" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, c.ID)
	return err
}
