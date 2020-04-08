package usecase

import (
	"strconv"

	"github.com/kwantz/simple-crud/configs"
	"github.com/kwantz/simple-crud/internal/app/repo"
)

// DeleteProductData : controller for ProductDeleteHandler
func DeleteProductData(productID int64) error {
	if _, err := repo.DeleteProduct(productID); err != nil {
		return err
	}

	productKeyRedis := "product-" + strconv.FormatInt(productID, 10)
	if err := configs.RedisClient.Del(productKeyRedis).Err(); err != nil {
		return err
	}

	return nil
}
