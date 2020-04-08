package usecase

import (
	"github.com/kwantz/simple-crud/internal/app/repo"
	"github.com/kwantz/simple-crud/internal/entity"
)

// GetProductListData : controller for ProductListGetHandler
func GetProductListData() (interface{}, error) {
	products := entity.ProductList{}
	productRepository, err := repo.FindProduct()
	if err != nil {
		return nil, err
	}

	for productRepository.Next() {
		product := entity.Product{}
		err := productRepository.Scan(
			&product.ID,
			&product.ShopID,
			&product.CategoryID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}
		products.List = append(products.List, product)
	}

	return products, nil
}
