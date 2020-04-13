package handler

import "net/http"

// ProductHandler interface
type ProductHandler interface {
	GetProductHandler(r *http.Request) []byte
	GetProductListHandler(r *http.Request) []byte

	PutProductHandler(r *http.Request) []byte
	PostProductHandler(r *http.Request) []byte
	DeleteProductHandler(r *http.Request) []byte
}
