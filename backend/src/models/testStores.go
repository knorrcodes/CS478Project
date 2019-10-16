package models

// TestProductStore builds the structure
type TestProductStore struct{}

// Save function for TestProductStore
func (s *TestProductStore) Save(p *Product) error { return nil }

// Delete function for TestProductStore
func (s *TestProductStore) Delete(p *Product) error { return nil }
