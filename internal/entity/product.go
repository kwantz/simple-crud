package entity

// Product entity
type Product struct {
	ID          int64  `json:"id"`
	ShopID      int64  `json:"shop_id"`
	CategoryID  int64  `json:"category_id"`
	Name        string `json:"name"`
	Price       int32  `json:"price"`
	Stock       int32  `json:"stock"`
	Description string `json:"description"`
}

// ProductList entity
type ProductList struct {
	List []Product `json:"list"`
}
