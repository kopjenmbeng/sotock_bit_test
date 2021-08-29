package movies

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	CtxAuthseCaseKey = "auth_usecase"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(InjectUseCaseContext)
		r.Get("/search", SearchHandler)
		r.Post("/detail", GetDetailHandler)
	})
	return r
}

func InjectUseCaseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo := NewMovieRepository()
		usecase := NewMoviesUsecase(repo)
		ctx := context.WithValue(r.Context(), CtxAuthseCaseKey, usecase)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UseCaseFromContext(c context.Context) IAuthenticationUseCase {
	return c.Value(CtxAuthseCaseKey).(IAuthenticationUseCase)
}