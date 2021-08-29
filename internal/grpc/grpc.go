package grpc

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	logrusLogger *log.Logger
	customFunc   grpc_logrus.CodeToLevel
)

type Server struct {
	log     *log.Logger
	address string
	srv     *grpc.Server
}

func NewServer(log *log.Logger, address string) *Server {
	// newrelic.newcon
	s := &Server{
		log:     log,
		address: address,
	}

	return s
}


func (s *Server) Serve() {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		s.log.Fatalln("Failed to listen:", err)
	}
	logrusLogger = s.log
	logrusEntry := logrus.NewEntry(logrusLogger)
	customFunc = func(code codes.Code) log.Level {
		if code == codes.OK {
			return log.InfoLevel
		}
		return log.ErrorLevel
	}

	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(customFunc),
	}

	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	// Create a server, make sure we put the grpc_ctxtags context before everything else.
	s.srv = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
		),
	)

	s.Register()
	s.log.Printf("start grpc server at port : %s", s.address)
	err = s.srv.Serve(lis)
	if err != nil {
		s.log.Fatalln("failed to serve grpc", err)
		// log.Fatalln(" failed to serve grpc")
	}

}
