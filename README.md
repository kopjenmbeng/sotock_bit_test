# sotock_bit_test

#install proto gen lib
- go get -u github.com/golang/protobuf/protoc-gen-go
- go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

#generate proto with plugin
- protoc --go_out=plugins=grpc:.  internal/model/receptionist.proto 

#generate proto not plugin 
- - protoc --go_out=:. --go_opt paths=source_relative  --go-grpc_out=:. --go-grpc_opt paths=source_relative  internal/model/*.proto 


#run grpc server
- go run main.go grpc

