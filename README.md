# sotock_bit_test
# task number 1
- this task is simple query ![Please click for view the result](../master/script/simple_database_query.sql)
# task number 2
- this task is create microservice from http://www.omdbapi.com/

    # install proto gen lib
    - go get -u github.com/golang/protobuf/protoc-gen-go
    - go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

    # generate proto with plugin
    - protoc --go_out=plugins=grpc:.  internal/proto/movies.proto 

    # generate proto not plugin 
    - - protoc --go_out=:. --go_opt paths=source_relative  --go-grpc_out=:. --go-grpc_opt paths=source_relative  internal/proto/*.proto 


    # run grpc server
    - go run main.go grpc
    - the service wil run at port 8080

    # run api server
    - go run main.go api
    - the grpc server wil run at port 8090

    # run unit test
    - go test -v .\test\

# task number 3
- this task is refactore code from funtion findFirstStringInBracket please execute command go run .\logic_task\findstring\ or you can find at logic_task/findstring for find the code

# task number 4
- this task is grouping anagram please execute command go run .\logic_task\anagram\ or you can find at logic_task/anagram for find the code
