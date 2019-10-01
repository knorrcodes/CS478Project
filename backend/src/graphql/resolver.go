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
