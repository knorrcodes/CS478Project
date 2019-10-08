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

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Product.GetProducts()
}

func (r *queryResolver) Category(ctx context.Context, id int) (*models.Category, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Category.GetCategoryByID(id)
}

func (r *queryResolver) Categories(ctx context.Context) ([]*models.Category, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Category.GetCategories()
}

func (r *queryResolver) Server(ctx context.Context, code int) (*models.Server, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Server.GetByCode(code)
}

func (r *queryResolver) Table(ctx context.Context, id int) (*models.Table, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Table.GetTableByID(id)
}

func (r *queryResolver) Tables(ctx context.Context) ([]*models.Table, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Table.GetTables()
}

func (r *queryResolver) Custcode(ctx context.Context, id *int, code *string) (*models.CustCode, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	if id != nil && *id > 0 {
		return storeCollection.CustCode.GetCustCodeByID(*id)
	}

	if code != nil && *code != "" {
		return storeCollection.CustCode.GetCustCodeByCode(*code)
	}

	return nil, nil
}

func (r *queryResolver) Custcodes(ctx context.Context) ([]*models.CustCode, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.CustCode.GetCustCodes()
}

func (r *queryResolver) Orders(ctx context.Context, server *int, status *OrderStatus) ([]*models.Order, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	if server != nil && *server > 0 {
		return r.ordersByServer(storeCollection.Order, *server, status)
	}

	return storeCollection.Order.GetOrders(graphQLOrderStatusToStore(status))
}

func (r *queryResolver) ordersByServer(s stores.OrderStore, server int, status *OrderStatus) ([]*models.Order, error) {
	return s.GetOrdersByServer(server, graphQLOrderStatusToStore(status))
}

func graphQLOrderStatusToStore(status *OrderStatus) stores.OrderStatus {
	switch *status {
	case OrderStatusOpened:
		return stores.OrderStatusOpened
	case OrderStatusClosed:
		return stores.OrderStatusClosed
	default:
		return stores.OrderStatusAny
	}
}

func (r *queryResolver) Order(ctx context.Context, id *int, table *int) (*models.Order, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	if id != nil && *id > 0 {
		return storeCollection.Order.GetOrderByID(*id)
	}

	if table != nil && *table > 0 {
		return storeCollection.Order.GetLatestOrderByTable(*table)
	}

	return nil, nil
}
