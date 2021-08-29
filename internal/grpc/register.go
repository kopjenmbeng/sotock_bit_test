package grpc

import (
	"github.com/jmoiron/sqlx"
	"github.com/kopjenmbeng/sotock_bit_test/internal/application/movies"
)

func (s *Server) Register(dbr sqlx.QueryerContext, dbw *sqlx.DB) {
	movies.NewMoviesGrpcHandler(s.log, s.srv,dbr,dbw)
}
