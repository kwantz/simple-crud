package usecase

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/kwantz/simple-crud/configs"
	"github.com/kwantz/simple-crud/internal/app/repo"
	"github.com/kwantz/simple-crud/internal/entity"
)

// GetSingleProductData : controller for ProductGetHandler
func GetSingleProductData(productID int64) (interface{}, error) {
	productKeyRedis := "product-" + strconv.FormatInt(productID, 10)
	productFromRedis, err := getProductFromRedis(productKeyRedis)
	if err != nil {
		return nil, err
	} else if productFromRedis != nil {
		return productFromRedis, nil
	}

	productFromMysql, err := getProductFromMysql(productID)
	if err != nil {
		return nil, err
	}

	if err := setProductToRedis(productKeyRedis, productFromMysql); err != nil {
		return nil, err
	}

	return productFromMysql, nil
}

func getProductFromRedis(productKeyRedis string) (interface{}, error) {
	serializedValue, err := configs.RedisClient.Get(productKeyRedis).Result()
	if err == redis.Nil { // Key not found
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var deserializedValue interface{}
	if err := json.Unmarshal([]byte(serializedValue), &deserializedValue); err != nil {
		return nil, err
	}

	log.Println("Get product from Redis ...")
	return deserializedValue, nil
}

func getProductFromMysql(productID int64) (interface{}, error) {
	product := entity.Product{}
	err := repo.FindProductByID(productID).Scan(
		&product.ID,
		&product.ShopID,
		&product.CategoryID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	log.Println("Get product from MySQL ...")
	return product, nil
}

func setProductToRedis(productKeyRedis string, product interface{}) error {
	serializedValue, err := json.Marshal(product)
	if err != nil {
		return err
	}

	duration := time.Minute
	return configs.RedisClient.Set(productKeyRedis, string(serializedValue), duration).Err()
}
