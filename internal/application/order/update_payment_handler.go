package order

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/go-chi/render"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/respond"
)

func UpdatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var (
		// err error
		rc  = r.Context()
		req PaidOrderRequest
		err error
	)
	r.Body = http.MaxBytesReader(w, r.Body, 5*1024*1024)

	err = render.Bind(r, &req)
	if err != nil {
		middleware.GetLogEntry(r).Error(xerrs.Details(err, respond.ErrMaxStack))
		respond.Nay(w, r, http.StatusBadRequest, err)
		return
	}
	// app_code:=r.Header.Get("X-Client-id")

	useCase := UseCaseFromContext(rc)
	code, err := useCase.UpdatePayment(rc, req)
	if err != nil {
		middleware.GetLogEntry(r).Error(xerrs.Details(err, respond.ErrMaxStack))
		respond.Nay(w, r, code, err)
		return
	}
	respond.Yay(w, r, code, map[string]string{
		"message": "pembayaran berhasil !",
	})
	// return

}
