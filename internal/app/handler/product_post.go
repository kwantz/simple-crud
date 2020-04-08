package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kwantz/simple-crud/internal/app/usecase"
	"github.com/kwantz/simple-crud/internal/entity"
)

// ProductPostHandler for route /products POST method
func ProductPostHandler(w http.ResponseWriter, r *http.Request) {
	data := &response{}
	dataBody := entity.Product{}

	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		data.Code = "400"
		data.Success = false
		data.Message = err.Error()

		w.Write(data.createJSON())
		return
	}

	product, err := usecase.AddProductData(dataBody)
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
	data.Data = product

	w.Write(data.createJSON())
	return
}
