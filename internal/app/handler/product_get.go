package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kwantz/simple-crud/internal/app/usecase"
)

// ProductListGetHandler for route /products GET method
func ProductListGetHandler(w http.ResponseWriter, r *http.Request) {
	data := &response{}
	productList, err := usecase.GetProductListData()
	if err != nil {
		data.Code = "500"
		data.Success = false
		data.Message = err.Error()

		w.Write(data.createJSON())
		return
	}

	data.Success = true
	data.Code = "200"
	data.Message = "Success"
	data.Data = productList

	w.Write(data.createJSON())
	return
}

// ProductGetHandler for route /products/{productID} GET method
func ProductGetHandler(w http.ResponseWriter, r *http.Request) {
	data := &response{}
	productID, err := strconv.ParseInt(chi.URLParam(r, "productID"), 10, 64)
	if err != nil {
		data.Code = "404"
		data.Success = false
		data.Message = "Not found"

		w.Write(data.createJSON())
		return
	}

	product, err := usecase.GetSingleProductData(productID)
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
