package entity

// Product entity
type Product struct {
	ID          int64  `json:"id"          db:"id"`
	ShopID      int64  `json:"shop_id"     db:"shop_id"`
	CategoryID  int64  `json:"category_id" db:"category_id"`
	Name        string `json:"name"        db:"name"`
	Price       int32  `json:"price"       db:"price"`
	Stock       int32  `json:"stock"       db:"stock"`
	Description string `json:"description" db:"description"`
}

// ProductList entity
type ProductList struct {
	List []Product `json:"list"`
}
