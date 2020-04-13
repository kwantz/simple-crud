package repository

import (
	"github.com/kwantz/simple-crud/internal/entity"
)

// ProductRepository interface
type ProductRepository interface {
	FindAll() (*entity.ProductList, error)
	Find(id int64) (*entity.Product, error)
	Store(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id int64) error

	FindInCache(id int64) (*entity.Product, error)
	StoreInCache(id int64, product *entity.Product) error
	DeleteInCache(id int64) error
}
