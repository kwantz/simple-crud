package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	uc "github.com/kwantz/simple-crud/internal/app/usecase"
	"github.com/kwantz/simple-crud/internal/entity"
)

type productHandler struct {
	usecase uc.ProductUsecase
}

// NewProductHandler - Create ProductHandler with inject ProductUsecase
func NewProductHandler(usecase uc.ProductUsecase) ProductHandler {
	return productHandler{usecase}
}

func (handler productHandler) GetProductHandler(r *http.Request) []byte {
	productID, err := strconv.ParseInt(chi.URLParam(r, "productID"), 10, 64)
	if err != nil {
		return entity.SetNotFoundResponse()
	}
	product, err := handler.usecase.FindThenStoreInCache(productID)
	if err != nil {
		return entity.SetServerErrorResponse(err.Error())
	}
	return entity.SetSuccessResponse(product)
}

func (handler productHandler) GetProductListHandler(r *http.Request) []byte {
	productList, err := handler.usecase.FindAll()
	if err != nil {
		return entity.SetServerErrorResponse(err.Error())
	}
	return entity.SetSuccessResponse(productList)
}

func (handler productHandler) PutProductHandler(r *http.Request) []byte {
	product := entity.Product{}
	productID, err := strconv.ParseInt(chi.URLParam(r, "productID"), 10, 64)
	if err != nil {
		return entity.SetNotFoundResponse()
	}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return entity.SetClientErrorResponse(err.Error())
	}
	if err = handler.usecase.UpdateThenDeleteInCache(productID, &product); err != nil {
		return entity.SetServerErrorResponse(err.Error())
	}
	return entity.SetSuccessResponse(product)
}

func (handler productHandler) PostProductHandler(r *http.Request) []byte {
	product := entity.Product{}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return entity.SetClientErrorResponse(err.Error())
	}
	if err := handler.usecase.Store(&product); err != nil {
		return entity.SetServerErrorResponse(err.Error())
	}
	return entity.SetCreatedResponse(product)
}

func (handler productHandler) DeleteProductHandler(r *http.Request) []byte {
	productID, err := strconv.ParseInt(chi.URLParam(r, "productID"), 10, 64)
	if err != nil {
		return entity.SetNotFoundResponse()
	}
	if err := handler.usecase.Delete(productID); err != nil {
		return entity.SetServerErrorResponse(err.Error())
	}
	return entity.SetSuccessResponse(nil)
}
