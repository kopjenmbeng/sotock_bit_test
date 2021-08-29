package products

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware/db_context"
)

const (
	CtxProductseCaseKey = "product_usecase"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(InjectUseCaseContext)
		r.Post("/", AddProductHandler)
		r.Put("/", UpdateProductHandler)
		r.Get("/", GetProductHandler)
		r.Delete("/{product_id}", DeleteProductHandler)
	})
	return r
}

func InjectUseCaseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbr := db_context.GetDbRead(r)
		dbw := db_context.GetDbWrite(r)
		repo := NewProductsRepository(dbr, dbw)
		usecase := NewProductsUseCase(repo)
		ctx := context.WithValue(r.Context(), CtxProductseCaseKey, usecase)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UseCaseFromContext(c context.Context) IProductsUsecase {
	return c.Value(CtxProductseCaseKey).(IProductsUsecase)
}
