package order

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/respond"
)

func GetMyOrderHandler(w http.ResponseWriter, r *http.Request) {
	var (
		// err error
		rc  = r.Context()
		err error
	)

	useCase := UseCaseFromContext(rc)
	data, code, err := useCase.GetMyOrder(rc)
	if err != nil {
		middleware.GetLogEntry(r).Error(xerrs.Details(err, respond.ErrMaxStack))
		respond.Nay(w, r, code, err)
		return
	}
	respond.Yay(w, r, code, data)
	return

}
