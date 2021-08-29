package order

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware/db_context"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware/jwe_auth"
)

const (
	CtxOrderCaseKey = "order_usecase"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(jwe_auth.GuardAnonymous(jwe_auth.TokenFromHeader))
		r.Use(InjectUseCaseContext)
		r.Post("/add", CreateOrderHandler)
		r.Put("/pay", UpdatePaymentHandler)
		r.Get("/my_order", GetMyOrderHandler)
		// r.Delete("/delete/{chart_id}",DeleteChartHandler)
	})
	return r
}

func InjectUseCaseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbr := db_context.GetDbRead(r)
		dbw := db_context.GetDbWrite(r)
		repo := NewOrderRepository(dbr, dbw)
		usecase := NewOrderUserCase(repo, r)
		ctx := context.WithValue(r.Context(), CtxOrderCaseKey, usecase)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UseCaseFromContext(c context.Context) IOrderUseCase {
	return c.Value(CtxOrderCaseKey).(IOrderUseCase)
}
