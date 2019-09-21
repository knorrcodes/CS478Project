package controllers

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"koala.pos/src/common"
	"koala.pos/src/models/stores"
)

type Product struct {
	e     *common.Environment
	store stores.ProductStore
}

func NewProductController(e *common.Environment, ps stores.ProductStore) *Product {
	return &Product{
		e:     e,
		store: ps,
	}
}

func (d *Product) GetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]interface{}{
		"time": time.Now().Format(time.RFC3339),
	}

	common.NewAPIResponse("Here you go", data).WriteResponse(w, http.StatusOK)
}
