package grpc

import (
	"github.com/jmoiron/sqlx"
	"github.com/kopjenmbeng/sotock_bit_test/internal/application/movies"
)

func (s *Server) Register(dbr sqlx.QueryerContext, dbw *sqlx.DB) {
	// register all rpc service
	movies.NewMoviesGrpcHandler(s.log, s.srv,dbr,dbw)
}
