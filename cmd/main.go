package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi"
	"github.com/kwantz/simple-crud/configs"

	"github.com/kwantz/simple-crud/internal/app/handler"
	"github.com/kwantz/simple-crud/internal/app/repository"
	"github.com/kwantz/simple-crud/internal/app/usecase"
)

func main() {
	mysql := configs.ConnectMySQL()
	defer mysql.Close()

	redis := configs.ConnectRedis()
	defer redis.Close()

	productRepository := repository.NewProductRepository(mysql, redis)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	router := chi.NewRouter()
	router.Post("/products", writeHandler(productHandler.PostProductHandler))
	router.Get("/products", writeHandler(productHandler.GetProductListHandler))
	router.Get("/products/{productID}", writeHandler(productHandler.GetProductHandler))
	router.Put("/products/{productID}", writeHandler(productHandler.PutProductHandler))
	router.Delete("/products/{productID}", writeHandler(productHandler.DeleteProductHandler))

	log.Println("Serving...")
	http.ListenAndServe(":8000", router)
}

type readHandler func(*http.Request) []byte

func writeHandler(handler readHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(handler(r))
	}
}
