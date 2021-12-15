# Servers microservices backend

### Implements microservice of [product](https://github.com/isaqueveras/products-microservice)
> Product microservice built in golang and gRPC

### Constants of system
```go
const (
  // port of server
  PortServer = "localhost:8080"
  // url of product service
  ProductURL = "localhost:50051"
  // Timeout to context
  ContextWithTimeout = time.Second * 10
)
```

### Command to run the test
```go
$ go test ./... --cover
```
