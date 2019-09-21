package models

type TestProductStore struct{}

func (s *TestProductStore) Save(p *Product) error   { return nil }
func (s *TestProductStore) Delete(p *Product) error { return nil }
