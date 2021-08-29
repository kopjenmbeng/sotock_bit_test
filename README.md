# stock bit test
# task number 1
- this task is simple query ![Please click for view the result](../master/script/simple_database_query.sql)
# task number 2
- this task is create microservice from http://www.omdbapi.com/

    # Technology Stack
-   GO
-   mysql
-   Consul for store configuration (Optional)
-   Postman for API Documentation

    # Before To RUN


    -   makesure .env file is exist and the configuration is like this.
        # predefined goconf env vars
        - GOCONF_ENV_PREFIX=stockbit
        - #GOCONF_CONSUL=localhost:8500 (please remark using # if you don't have consul so it will read ![this json config](../master/stockbit.config.json))
        - GOCONF_TYPE=json
        - GOCONF_FILENAME=stockbit.config

        # Newrelic
        - #PROPERTY_NEWRELIC_KEY=

    After you full fill the requeirment above this is rule for run application locally.
    # run grpc server
    - go run main.go grpc
    - the service wil run at port 8080

    # run api server
    - go run main.go api
    - the grpc server wil run at port 8090

    # run unit test
    - go test -v .\test\

    # api documentation
    - for Api documentation please import ![api doc](../master/document/stockbit.postman_collection.json))
    

    

# task number 3
- this task is refactore code from funtion findFirstStringInBracket please execute command go run .\logic_task\findstring\ or you can find at logic_task/findstring for find the code

# task number 4
- this task is grouping anagram please execute command go run .\logic_task\anagram\ or you can find at logic_task/anagram for find the code
