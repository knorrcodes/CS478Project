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

func (r *Resolver) CustCode() CustCodeResolver {
	return &custCodeResolver{r}
}

// Product is the product resolver
func (r *Resolver) Product() ProductResolver {
	return &productResolver{r}
}

func (r *Resolver) Order() OrderResolver {
	return &orderResolver{r}
}

func (r *Resolver) Table() TableResolver {
	return &tableResolver{r}
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

type custCodeResolver struct{ *Resolver }

func (r *custCodeResolver) Order(ctx context.Context, obj *models.CustCode) (*models.Order, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Order.GetOrderByID(obj.OrderID)
}

type orderResolver struct{ *Resolver }

func (r *orderResolver) Table(ctx context.Context, obj *models.Order) (*models.Table, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Table.GetTableByID(obj.TableID)
}

func (r *orderResolver) Server(ctx context.Context, obj *models.Order) (*models.Server, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Server.GetByID(obj.ServerID)
}

type tableResolver struct{ *Resolver }

func (r *tableResolver) Orders(ctx context.Context, obj *models.Table) ([]*models.Order, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Order.GetOrdersByTable(obj.ID, stores.OrderStatusAny)
}
