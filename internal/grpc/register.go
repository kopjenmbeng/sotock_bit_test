package grpc

import(
	"github.com/kopjenmbeng/sotock_bit_test/internal/application/movies"
)
func (s *Server) Register() {
	movies.NewMoviesGrpcHandler(s.log,s.srv)
}