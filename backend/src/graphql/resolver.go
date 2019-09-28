package graphql

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"koala.pos/src/models"
	"koala.pos/src/models/stores"
)

// Resolver is the root GraphQL resolver
type Resolver struct{}

// Product is the product resolver
func (r *Resolver) Product() ProductResolver {
	return &productResolver{r}
}

// Query is the query resolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Mutation is the mutation resolver
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type productResolver struct{ *Resolver }

func (r *productResolver) Category(ctx context.Context, obj *models.Product) (*models.Category, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Category.GetCategoryByID(obj.Category)
}

type queryResolver struct{ *Resolver }

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

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateProduct(ctx context.Context, input NewProduct) (*models.Product, error) {
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
