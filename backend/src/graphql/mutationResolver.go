package graphql

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"errors"
	"time"

	"koala.pos/src/auth"
	"koala.pos/src/common"
	"koala.pos/src/models"
	"koala.pos/src/models/stores"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateProduct(ctx context.Context, input NewProductInput) (*models.Product, error) {
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
	product.Category = input.Category
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

func (r *mutationResolver) CreateCategory(ctx context.Context, name string) (*models.Category, error) {
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
	cat.Name = name
	if err := cat.Save(); err != nil {
		return nil, err
	}

	return cat, nil
}

func (r *mutationResolver) CreateTable(ctx context.Context, num int) (*models.Table, error) {
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
	table.Num = num
	if err := table.Save(); err != nil {
		return nil, err
	}

	return table, nil
}

func (r *mutationResolver) CreateCustCode(ctx context.Context, orderID int) (*models.CustCode, error) {
	// Any server can create a customer code

	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	code := models.NewCustCode(storeCollection.CustCode)
	code.StartTime = time.Now()
	code.EndTime = code.StartTime.Add(2 * time.Hour)
	code.Code = common.GenerateOneTimeCode()
	code.OrderID = orderID

	if err := code.Save(); err != nil {
		return nil, err
	}

	return code, nil
}

func (r *mutationResolver) StartOrder(ctx context.Context, input NewOrderInput) (*models.Order, error) {
	// Any server can create a customer code

	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	server := auth.GetServerFromContext(ctx)

	order := models.NewOrder(storeCollection.Order)
	order.StartTime = time.Now()
	order.EndTime = time.Unix(0, 0)
	order.TableID = input.Table
	order.ServerID = server.ID

	if err := order.Save(); err != nil {
		return nil, err
	}

	return order, nil
}

func (r *mutationResolver) CloseOrder(ctx context.Context, id int) (*models.Order, error) {
	// Any server can create a customer code

	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	order, err := storeCollection.Order.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	order.EndTime = time.Now()
	if err := order.Save(); err != nil {
		return nil, err
	}

	return order, nil
}

func (r *mutationResolver) AddItemToOrder(ctx context.Context, orderID int, products []int) (*models.OrderItem, error) {
	// Any server can update an order

	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	order, err := storeCollection.Order.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	if !order.IsOpen() {
		return nil, errors.New("cannot add items to a closed order")
	}

	orderItem := models.NewOrderItem(storeCollection.OrderItem)
	orderItem.OrderID = orderID
	orderItem.Products = products
	if err := orderItem.Save(); err != nil {
		return nil, err
	}

	return orderItem, nil
}

func (r *mutationResolver) ApplyPayment(ctx context.Context, input AddPaymentInput) (*models.Payment, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	order, err := storeCollection.Order.GetOrderByID(input.Order)
	if err != nil {
		return nil, err
	}

	if !order.IsOpen() {
		return nil, errors.New("cannot add payment to a closed order")
	}

	payment := models.NewPayment(storeCollection.Payment)
	payment.OrderID = input.Order
	payment.Amount = input.Amount
	payment.Timestamp = time.Now()
	if err := payment.Save(); err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *mutationResolver) DeleteOrderItem(ctx context.Context, id int) (*models.Order, error) {
	storeCollection := stores.GetStoreCollectionFromContext(ctx)
	if storeCollection == nil {
		return nil, errors.New("Failed to get storage")
	}

	orderItem, err := storeCollection.OrderItem.GetByID(id)
	if err != nil {
		return nil, err
	}

	order, err := storeCollection.Order.GetOrderByID(orderItem.OrderID)
	if err != nil {
		return nil, err
	}

	if err := orderItem.Delete(); err != nil {
		return nil, err
	}
	return order, nil
}
