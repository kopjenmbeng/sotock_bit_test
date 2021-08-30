package movies

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware/db_context"
)

const (
	CtxMoviesCaseKey = "movies_usecase"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(InjectUseCaseContext)
		r.Get("/search", SearchHandler)
		r.Get("/get_detail/{id}", GetDetailHandler)
	})
	return r
}

// declare dependency injection
func InjectUseCaseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbr := db_context.GetDbRead(r)
		dbw := db_context.GetDbWrite(r)
		repo := NewMovieRepository(dbr, dbw)
		usecase := NewMoviesUsecase(repo)
		ctx := context.WithValue(r.Context(), CtxMoviesCaseKey, usecase)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UseCaseFromContext(c context.Context) IMoviesUsecase {
	return c.Value(CtxMoviesCaseKey).(IMoviesUsecase)
}
