package graphql

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"koala.pos/src/auth"
	"koala.pos/src/models"
	"koala.pos/src/models/stores"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateProduct(ctx context.Context, input NewProduct) (*models.Product, error) {
	server := auth.GetServerFromContext(ctx)
	if server == nil {
		return nil, errors.New("Failed to check permissions")
	}

	if !server.Manager {
		return nil, errors.New("insufficient privilages")
	}

	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	product := models.NewProduct(storeCollection.Product)

	// Required fields
	product.Name = input.Name
	product.Price = input.Price
	product.Category = input.Category.ID
	product.WSCost = input.Wscost

	// Optional fields
	if input.Desc != nil {
		product.Desc = *input.Desc
	}
	if input.Picture != nil {
		product.Picture = *input.Picture
	}
	if input.NumOfSides != nil {
		product.NumOfSides = *input.NumOfSides
	}

	if err := product.Save(); err != nil {
		return nil, err
	}

	return product, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input NewCategory) (*models.Category, error) {
	server := auth.GetServerFromContext(ctx)
	if server == nil {
		return nil, errors.New("Failed to check permissions")
	}

	if !server.Manager {
		return nil, errors.New("insufficient privilages")
	}

	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	cat := models.NewCategory(storeCollection.Category)
	cat.Name = input.Name
	if err := cat.Save(); err != nil {
		return nil, err
	}

	return cat, nil
}

func (r *mutationResolver) CreateTable(ctx context.Context, input NewTable) (*models.Table, error) {
	server := auth.GetServerFromContext(ctx)
	if server == nil {
		return nil, errors.New("Failed to check permissions")
	}

	if !server.Manager {
		return nil, errors.New("insufficient privilages")
	}

	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	table := models.NewTable(storeCollection.Table)
	table.Num = input.Num
	if err := table.Save(); err != nil {
		return nil, err
	}

	return table, nil
}
