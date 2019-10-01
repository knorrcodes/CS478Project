package graphql

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"koala.pos/src/models"
	"koala.pos/src/models/stores"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Product(ctx context.Context, id int) (*models.Product, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Product.GetProductByID(id)
}

func (r *queryResolver) Category(ctx context.Context, id int) (*models.Category, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Category.GetCategoryByID(id)
}

func (r *queryResolver) Server(ctx context.Context, code int) (*models.Server, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Server.GetByCode(code)
}

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Product.GetProducts()
}

func (r *queryResolver) Categories(ctx context.Context) ([]*models.Category, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Category.GetCategories()
}
