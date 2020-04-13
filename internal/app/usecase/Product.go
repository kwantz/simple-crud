package usecase

import (
	"database/sql"

	"github.com/gomodule/redigo/redis"
	repo "github.com/kwantz/simple-crud/internal/app/repository"
	"github.com/kwantz/simple-crud/internal/entity"
)

type productUsecase struct {
	repository repo.ProductRepository
}

// NewProductUsecase - Create ProductUsecase with inject ProductRepository
func NewProductUsecase(repository repo.ProductRepository) ProductUsecase {
	return productUsecase{repository}
}

func (usecase productUsecase) FindInCache(id int64) (*entity.Product, error) {
	product, err := usecase.repository.FindInCache(id)
	if err == redis.ErrNil {
		return nil, nil
	}
	return product, err
}

func (usecase productUsecase) StoreInCache(id int64, product *entity.Product) error {
	err := usecase.repository.StoreInCache(id, product)
	return err
}

func (usecase productUsecase) DeleteInCache(id int64) error {
	err := usecase.repository.DeleteInCache(id)
	return err
}

func (usecase productUsecase) FindAll() (*entity.ProductList, error) {
	productList, err := usecase.repository.FindAll()
	return productList, err
}

func (usecase productUsecase) Find(id int64) (*entity.Product, error) {
	product, err := usecase.repository.Find(id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return product, err
}

func (usecase productUsecase) Store(product *entity.Product) error {
	err := usecase.repository.Store(product)
	return err
}

func (usecase productUsecase) Update(id int64, product *entity.Product) error {
	product.ID = id
	err := usecase.repository.Update(product)
	return err
}

func (usecase productUsecase) Delete(id int64) error {
	err := usecase.repository.Delete(id)
	if err == nil {
		err = usecase.DeleteInCache(id)
	}
	return err
}

func (usecase productUsecase) FindThenStoreInCache(id int64) (*entity.Product, error) {
	product, err := usecase.Find(id)
	if product != nil {
		err = usecase.StoreInCache(id, product)
	}
	return product, err
}

func (usecase productUsecase) UpdateThenDeleteInCache(id int64, product *entity.Product) error {
	err := usecase.Update(id, product)
	if err == nil {
		err = usecase.DeleteInCache(id)
	}
	return err
}
