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

// CustCode is the customer code resolver
func (r *Resolver) CustCode() CustCodeResolver {
	return &custCodeResolver{r}
}

// Product is the product resolver
func (r *Resolver) Product() ProductResolver {
	return &productResolver{r}
}

// Order is the order resolver
func (r *Resolver) Order() OrderResolver {
	return &orderResolver{r}
}

// OrderItem is the orderitem resolver
func (r *Resolver) OrderItem() OrderItemResolver {
	return &orderItemResolver{r}
}

// Payment is the payment resolver
func (r *Resolver) Payment() PaymentResolver {
	return &paymentResolver{r}
}

// Table is the table resolver
func (r *Resolver) Table() TableResolver {
	return &tableResolver{r}
}

// Category is the category resolver
func (r *Resolver) Category() CategoryResolver {
	return &categoryResolver{r}
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

func (r *orderResolver) CustCode(ctx context.Context, obj *models.Order) (*models.CustCode, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.CustCode.GetCustCodeByOrderID(obj.ID)
}

func (r *orderResolver) Items(ctx context.Context, obj *models.Order) ([]*models.OrderItem, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.OrderItem.GetByOrder(obj.ID)
}

func (r *orderResolver) Payments(ctx context.Context, obj *models.Order) ([]*models.Payment, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Payment.GetPaymentsForOrder(obj.ID)
}

type tableResolver struct{ *Resolver }

func (r *tableResolver) Orders(ctx context.Context, obj *models.Table, status *OrderStatus) ([]*models.Order, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	ordStatus := stores.OrderStatusOpened
	if *status == OrderStatusClosed {
		ordStatus = stores.OrderStatusClosed
	} else if *status == OrderStatusAny {
		ordStatus = stores.OrderStatusAny
	}

	return storeCollection.Order.GetOrdersByTable(obj.ID, ordStatus)
}

type orderItemResolver struct{ *Resolver }

func (r *orderItemResolver) Products(ctx context.Context, obj *models.OrderItem) ([]*models.Product, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Product.GetProductsByID(obj.Products)
}

func (r *orderItemResolver) Order(ctx context.Context, obj *models.OrderItem) (*models.Order, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Order.GetOrderByID(obj.ID)
}

type paymentResolver struct{ *Resolver }

func (r *paymentResolver) Order(ctx context.Context, obj *models.Payment) (*models.Order, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Order.GetOrderByID(obj.OrderID)
}

type categoryResolver struct{ *Resolver }

func (r *categoryResolver) Products(ctx context.Context, obj *models.Category) ([]*models.Product, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	return storeCollection.Product.GetProductsByCategory(obj.ID)
}
