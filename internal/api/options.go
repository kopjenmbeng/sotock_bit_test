package api

import (
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware/jwe_auth"
)

func JWE(jw *jwe_auth.JWE) Option {
	return func(s *Server) {
		s.jwe = jw
	}
}
