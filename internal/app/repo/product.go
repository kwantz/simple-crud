package repo

import (
	"database/sql"

	"github.com/kwantz/simple-crud/configs"
	"github.com/kwantz/simple-crud/internal/entity"
)

// FindProduct : get all data
func FindProduct() (*sql.Rows, error) {
	return configs.MysqlClient.Query(`
		SELECT id, shop_id, category_id, name, price, stock, description
		FROM Product
	`)
}

// FindProductByID : get single data
func FindProductByID(productID int64) *sql.Row {
	return configs.MysqlClient.QueryRow(`
		SELECT id, shop_id, category_id, name, price, stock, description
		FROM Product
		WHERE id=?
	`,
		productID,
	)
}

// InsertProduct : store single data
func InsertProduct(productEntity entity.Product) (sql.Result, error) {
	return configs.MysqlClient.Exec(`
		INSERT INTO Product(shop_id, category_id, name, price, stock, description)
		VALUES(?, ?, ?, ?, ?, ?)
	`,
		productEntity.ShopID,
		productEntity.CategoryID,
		productEntity.Name,
		productEntity.Price,
		productEntity.Stock,
		productEntity.Description,
	)
}

// UpdateProduct : update single data
func UpdateProduct(productID int64, productEntity entity.Product) (sql.Result, error) {
	return configs.MysqlClient.Exec(`
		UPDATE Product SET category_id=?, name=?, price=?, stock=?, description=?
		WHERE id=?
	`,
		productEntity.CategoryID,
		productEntity.Name,
		productEntity.Price,
		productEntity.Stock,
		productEntity.Description,
		productID,
	)
}

// DeleteProduct : delete single data
func DeleteProduct(productID int64) (sql.Result, error) {
	return configs.MysqlClient.Exec(`
		DELETE FROM Product
		WHERE id=?
	`,
		productID,
	)
}
