package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kwantz/simple-crud/internal/app/usecase"
	"github.com/kwantz/simple-crud/internal/entity"
)

// ProductPutHandler for route /products/{productID} PUT method
func ProductPutHandler(w http.ResponseWriter, r *http.Request) {
	data := &response{}
	productEntity := entity.Product{}

	if err := json.NewDecoder(r.Body).Decode(&productEntity); err != nil {
		data.Code = "400"
		data.Success = false
		data.Message = err.Error()

		w.Write(data.createJSON())
		return
	}

	productID, err := strconv.ParseInt(chi.URLParam(r, "productID"), 10, 64)
	if err != nil {
		data.Code = "404"
		data.Success = false
		data.Message = "Not found"

		w.Write(data.createJSON())
		return
	}

	product, err := usecase.EditProductData(productID, productEntity)
	if err != nil {
		data.Code = "500"
		data.Success = false
		data.Message = err.Error()

		w.Write(data.createJSON())
		return
	}
	if product == nil {
		data.Code = "404"
		data.Success = false
		data.Message = "Not found"

		w.Write(data.createJSON())
		return
	}

	data.Success = true
	data.Code = "200"
	data.Message = "Success"
	data.Data = product

	w.Write(data.createJSON())
	return
}
