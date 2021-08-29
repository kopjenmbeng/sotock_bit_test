package movie

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	CtxMovieUseCaseKey = "movie_usecase"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(InjectUseCaseContext)
		// r.Get("/search", SearchHandler)
		// r.Post("/register", RegisterHandler)
	})
	return r
}

func InjectUseCaseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// dbr := db_context.GetDbRead(r)
		// dbw := db_context.GetDbWrite(r)
		repo := NewMovieRepository()
		usecase := NewMovieUsecase(repo)
		ctx := context.WithValue(r.Context(), CtxMovieUseCaseKey, usecase)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UseCaseFromContext(c context.Context) IMovieUsecase {
	return c.Value(CtxMovieUseCaseKey).(IMovieUsecase)
}
