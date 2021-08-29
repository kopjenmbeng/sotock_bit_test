package api

import (
	"github.com/go-chi/chi"
	// "github.com/kopjenmbeng/sotock_bit_test/internal/application/authentication"
	// "github.com/kopjenmbeng/sotock_bit_test/internal/application/order"
	"github.com/kopjenmbeng/sotock_bit_test/internal/application/movie"
	// "github.com/kopjenmbeng/sotock_bit_test/internal/application/tes"
)

func routes(r *chi.Mux) {
	r.Route("/v1", func(r chi.Router) {
		// r.Mount("/authentication", authentication.Routes())
		// r.Mount("/products", products.Routes())
		// r.Mount("/order", order.Routes())
		r.Mount("/movie",movie.Routes())
	})
}
