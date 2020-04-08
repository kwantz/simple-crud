package usecase

import (
	"strconv"

	"github.com/kwantz/simple-crud/configs"
	"github.com/kwantz/simple-crud/internal/app/repo"
	"github.com/kwantz/simple-crud/internal/entity"
)

// EditProductData : controller for ProductPutHandler
func EditProductData(productID int64, productEntity entity.Product) (interface{}, error) {
	productRepository, err := repo.UpdateProduct(productID, productEntity)
	if err != nil {
		return productEntity, err
	}

	isProductFound, err := productRepository.RowsAffected()
	if err != nil {
		return nil, err
	}
	if isProductFound == 0 {
		return nil, nil
	}

	productEntity.ID = productID
	productKeyRedis := "product-" + strconv.FormatInt(productID, 10)
	if err := configs.RedisClient.Del(productKeyRedis).Err(); err != nil {
		return nil, err
	}

	return productEntity, nil
}
