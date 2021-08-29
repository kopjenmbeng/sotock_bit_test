package products

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/respond"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var (
		// err error
		rc  = r.Context()
		err error
	)
	// app_code:=r.Header.Get("X-Client-id")
	// ProductId := chi.URLParam(r, "product_id")
	// if err := validator.ValidateEmpty("chart_id", ProductId); err != nil {
	// 	respond.Nay(w, r, http.StatusBadRequest, err)
	// 	return
	// }
	useCase := UseCaseFromContext(rc)
	data, code, err := useCase.GetAll(rc)
	if err != nil {
		middleware.GetLogEntry(r).Error(xerrs.Details(err, respond.ErrMaxStack))
		respond.Nay(w, r, code, err)
		return
	}
	respond.Yay(w, r, code, data)
	return

}
