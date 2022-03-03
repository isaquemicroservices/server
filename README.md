# Microservices Servers 

## (4) Microservices built with Golang, gRPC and PostgreSQL.
- [Product](https://github.com/isaquemicroservices/products)
- [Authentication](https://github.com/isaquemicroservices/authentication)
- [Customer](github.com/isaquemicroservices/customer)
- [Email](https://github.com/isaquemicroservices/email)

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
