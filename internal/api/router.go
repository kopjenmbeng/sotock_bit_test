package api

import (
	"github.com/go-chi/chi"
	"github.com/kopjenmbeng/sotock_bit_test/internal/application/movies"
)

func routes(r *chi.Mux) {
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/movie",movies.Routes())
	})
}
