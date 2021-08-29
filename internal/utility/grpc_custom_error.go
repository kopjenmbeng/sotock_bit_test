package utility

import (
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcInternalServerError(err error)error{
	return status.Error(codes.Code(code.Code_INTERNAL),err.Error())
}

func GrpcUnAuthenticateError(err error)error{
	return status.Error(codes.Code(code.Code_UNAUTHENTICATED),err.Error())
}