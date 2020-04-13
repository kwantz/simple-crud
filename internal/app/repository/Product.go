package repository

import (
	"encoding/json"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/kwantz/simple-crud/internal/entity"
)

type productRepository struct {
	db    *sqlx.DB
	cache redis.Conn
}

// NewProductRepository - Create ProductRepository with inject DB and cache
func NewProductRepository(db *sqlx.DB, cache redis.Conn) ProductRepository {
	return productRepository{db, cache}
}

func (repository productRepository) getProductKey(id int64) string {
	return "product-" + strconv.FormatInt(id, 10)
}

func (repository productRepository) FindAll() (*entity.ProductList, error) {
	query := `SELECT * FROM Product`

	productList := entity.ProductList{}
	err := repository.db.Select(&productList.List, query)
	return &productList, err
}

func (repository productRepository) Find(id int64) (*entity.Product, error) {
	product, err := repository.FindInCache(id)
	if err == redis.ErrNil {
		query := `SELECT * FROM Product WHERE id = ?`
		err = repository.db.Get(&product, query, id)
	}
	return product, err
}

func (repository productRepository) Store(product *entity.Product) error {
	query := `
		INSERT INTO Product(shop_id, category_id, name, price, stock, description)
		VALUES(:shop_id, :category_id, :name, :price, :stock, :description)
	`
	result, err := repository.db.NamedExec(query, product)
	if err == nil {
		product.ID, err = result.LastInsertId()
	}
	return err
}

func (repository productRepository) Update(product *entity.Product) error {
	query := `
		UPDATE Product SET
			category_id=:category_id, name=:name, price=:price,
			stock=:stock, description=:description
		WHERE id=:id
	`
	_, err := repository.db.NamedExec(query, product)
	return err
}

func (repository productRepository) Delete(id int64) error {
	query := `DELETE FROM Product WHERE id=?`

	_, err := repository.db.Exec(query, id)
	return err
}

func (repository productRepository) FindInCache(id int64) (*entity.Product, error) {
	value, err := redis.String(repository.cache.Do("GET", repository.getProductKey(id)))
	product := entity.Product{}
	if err == nil {
		err = json.Unmarshal([]byte(value), &product)
	}
	return &product, nil
}

func (repository productRepository) StoreInCache(id int64, product *entity.Product) error {
	value, err := json.Marshal(product)
	if err == nil {
		_, err = repository.cache.Do("SET", repository.getProductKey(id), string(value))
	}
	return err
}

func (repository productRepository) DeleteInCache(id int64) error {
	_, err := repository.cache.Do("DEL", repository.getProductKey(id))
	return err
}
