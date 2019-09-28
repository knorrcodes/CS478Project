package models

type ServerStore interface {
	Save(*Server) error
	Delete(*Server) error
}

type Server struct {
	store   ServerStore
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Code    int    `json:"code"`
	Manager bool   `json:"manager"`
}

// NewServer creates a new Server object
func NewServer(cs ServerStore) *Server {
	return &Server{
		store: cs,
	}
}

func (c *Server) IsNew() bool {
	return c.ID == 0
}

func (c *Server) Save() error {
	return c.store.Save(c)
}

func (c *Server) Delete() error {
	return c.store.Delete(c)
}
