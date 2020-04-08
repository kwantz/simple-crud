package usecase

import (
	"github.com/kwantz/simple-crud/internal/app/repo"
	"github.com/kwantz/simple-crud/internal/entity"
)

// AddProductData : controller for ProductPostHandler
func AddProductData(product entity.Product) (interface{}, error) {
	productRepository, err := repo.InsertProduct(product)
	if err != nil {
		return nil, err
	}

	id, err := productRepository.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.ID = id
	return product, nil
}
