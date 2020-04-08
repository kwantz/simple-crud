package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kwantz/simple-crud/configs"
	"github.com/kwantz/simple-crud/internal/app/handler"
)

func main() {
	r := chi.NewRouter()

	r.Get("/products", handler.ProductListGetHandler)
	r.Post("/products", handler.ProductPostHandler)

	r.Get("/products/{productID}", handler.ProductGetHandler)
	r.Put("/products/{productID}", handler.ProductPutHandler)
	r.Delete("/products/{productID}", handler.ProductDeleteHandler)

	configs.ConnectRedis()

	configs.ConnectMysql()
	defer configs.MysqlClient.Close()

	log.Println("Serving...")
	http.ListenAndServe(":8000", r)
}
