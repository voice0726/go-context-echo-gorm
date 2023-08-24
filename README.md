# go-context-echo-gorm

This is an experimental code to see if gorm chancel a db query when a request is closed.
To see:

1. Launch db container where MySQL runs
2. Run the server with `go run`
3. Open your browser and access to `http://localhost:8080/test`
4. After 5 sec., you will have "ok" response
5. Cancel your request within 5 sec., and see the log of the server. You will see gorm canceled a query when you canceled the request.
