package usecase

import "github.com/kwantz/simple-crud/internal/entity"

// ProductUsecase interface
type ProductUsecase interface {
	FindInCache(id int64) (*entity.Product, error)
	StoreInCache(id int64, product *entity.Product) error
	DeleteInCache(id int64) error

	FindAll() (*entity.ProductList, error)
	Find(id int64) (*entity.Product, error)
	Store(product *entity.Product) error
	Update(id int64, product *entity.Product) error
	Delete(id int64) error

	FindThenStoreInCache(id int64) (*entity.Product, error)
	UpdateThenDeleteInCache(id int64, product *entity.Product) error
}
