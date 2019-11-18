package models

// ServerStore interface type
type ServerStore interface {
	Save(*Server) error
	Delete(*Server) error
}

// Server object structure
type Server struct {
	store    ServerStore
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     int    `json:"code"`
	Manager  bool   `json:"manager"`
	CustCode string `json:"-"`
}

// NewServer creates a new Server object
func NewServer(cs ServerStore) *Server {
	return &Server{
		store: cs,
	}
}

// IsNew is the function to check if a server object is new
func (c *Server) IsNew() bool {
	return c.ID == 0
}

// Save is the function to save a product object
func (c *Server) Save() error {
	return c.store.Save(c)
}

// Delete is the function to delete a product object
func (c *Server) Delete() error {
	return c.store.Delete(c)
}
