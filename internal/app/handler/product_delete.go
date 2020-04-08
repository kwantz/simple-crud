package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kwantz/simple-crud/internal/app/usecase"
)

// ProductDeleteHandler for route /products/{productID} DELETE method
func ProductDeleteHandler(w http.ResponseWriter, r *http.Request) {
	data := &response{}

	productID, err := strconv.ParseInt(chi.URLParam(r, "productID"), 10, 64)
	if err != nil {
		data.Code = "400"
		data.Success = false
		data.Message = err.Error()

		w.Write(data.createJSON())
		return
	}

	if err := usecase.DeleteProductData(productID); err != nil {
		data.Code = "500"
		data.Success = false
		data.Message = err.Error()

		w.Write(data.createJSON())
		return
	}

	data.Success = true
	data.Code = "200"
	data.Message = "Success"

	w.Write(data.createJSON())
	return
}
