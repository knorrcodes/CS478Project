package stores

import (
	"errors"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

type ServerStore interface {
	Get() ([]*models.Server, error)
	GetByID(id int) (*models.Server, error)
	Save(p *models.Server) error
	Delete(p *models.Server) error
	GetByCode(code int) (*models.Server, error)
}

type Server struct {
	e *common.Environment
}

func NewServerStore(e *common.Environment) *Server {
	return &Server{
		e: e,
	}
}

func (s *Server) Get() ([]*models.Server, error) {
	return s.getFromDatabase("")
}

func (s *Server) GetByID(id int) (*models.Server, error) {
	if id == 0 {
		return nil, errors.New("Server ID required")
	}

	sql := `WHERE "id" = ?`
	Servers, err := s.getFromDatabase(sql, id)
	if len(Servers) == 0 {
		return nil, err
	}
	return Servers[0], err
}

func (s *Server) GetByCode(code int) (*models.Server, error) {
	if code == 0 {
		return nil, errors.New("Server code required")
	}

	sql := `WHERE "code" = ?`
	Servers, err := s.getFromDatabase(sql, code)
	if len(Servers) == 0 {
		return nil, err
	}
	return Servers[0], err
}

func (s *Server) getFromDatabase(where string, values ...interface{}) ([]*models.Server, error) {
	sql := `SELECT
				"id",
				"name",
				"code",
				"manager"
			FROM "server" ` + where

	rows, err := s.e.DB.Query(sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Server
	for rows.Next() {
		var id int
		var name string
		var code int
		var manager bool

		err := rows.Scan(
			&id,
			&name,
			&code,
			&manager,
		)
		if err != nil {
			continue
		}

		Server := models.NewServer(s)
		Server.ID = id
		Server.Name = name
		Server.Code = code
		Server.Manager = manager

		results = append(results, Server)
	}
	return results, nil
}

func (s *Server) Save(p *models.Server) error {
	if p.ID == 0 {
		return s.saveNew(p)
	}
	return s.updateExisting(p)
}

func (s *Server) updateExisting(p *models.Server) error {
	sql := `UPDATE "server"
			SET	"name" = ?,
				"code" = ?,
				"manager" = ?
			WHERE "id" = ?`

	_, err := s.e.DB.Exec(
		sql,
		p.Name,
		p.Code,
		p.Manager,
		p.ID,
	)
	return err
}

func (s *Server) saveNew(p *models.Server) error {
	if p.Name == "" {
		return errors.New("Server name cannot be empty")
	}

	sql := `INSERT INTO "server"
				("name",
				"code",
				"manager") VALUES (?,?,?)`

	result, err := s.e.DB.Exec(
		sql,
		p.Name,
		p.Code,
		p.Manager,
	)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)
	return nil
}

func (s *Server) Delete(p *models.Server) error {
	if p.ID == 0 {
		return nil
	}

	sql := `DELETE FROM "server" WHERE "id" = ?`
	_, err := s.e.DB.Exec(sql, p.ID)
	return err
}
