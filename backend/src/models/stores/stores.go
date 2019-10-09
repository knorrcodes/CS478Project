package stores

import (
	"context"
	"net/http"

	"koala.pos/src/common"
)

type StoreCollection struct {
	Product   ProductStore
	Category  CategoryStore
	Server    ServerStore
	Table     TableStore
	CustCode  CustCodeStore
	Order     OrderStore
	OrderItem OrderItemStore
}

// GetStoreCollectionFromContext retrieves the StoreCollection from the current request.
func GetStoreCollectionFromContext(ctx context.Context) *StoreCollection {
	if rv := ctx.Value(common.StoreCollectionKey); rv != nil {
		return rv.(*StoreCollection)
	}
	return nil
}

// SetStoreCollectionToContext sets an StoreCollection for the current request.
func SetStoreCollectionToContext(r *http.Request, e *StoreCollection) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), common.StoreCollectionKey, e))
}
