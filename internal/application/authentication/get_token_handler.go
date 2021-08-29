package authentication

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/respond"
)

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	var (
		// err error
		rc = r.Context()
	)

	// app_code:=r.Header.Get("X-Client-id")
	Email := r.URL.Query().Get("email")
	var password string = r.URL.Query().Get("password")
	useCase := UseCaseFromContext(rc)
	data, code, err := useCase.GetToken(rc, Email, password)
	if err != nil {
		middleware.GetLogEntry(r).Error(xerrs.Details(err, respond.ErrMaxStack))
		respond.Nay(w, r, code, err)
		return
	}
	respond.Yay(w, r, http.StatusOK, data)
	return

}
