# Servers microservices backend

### Implements microservice of [product](https://github.com/isaqueveras/products-microservice)
Product microservice build in golang and gRPC

### Implements microservice of [authentication](https://github.com/isaqueveras/authentication-microservice)
Authentication microservice build in golang and gRPC

### Create folder for config.json file
```bat
$ sudo mkdir /etc/server-ms-backend
$ sudo touch /etc/server-ms-backend/config.json
$ sudo cp ./config.json /etc/server-ms-backend/config.json
$ sudo chmod 777 /etc/server-ms-backend/config.json
```
if you changed the config.json file, use the command at the bottom to update the config.json file on your computer
```bat
$ sudo cp ./config.json /etc/server-ms-backend/config.json
```

### Command to run server
```bat
$ go run main.go
```

### Command to run the test
```bat
$ go test ./... --cover
```
