## Order Service

This service is built for Golang Beginner Course assignment.

### Architecture

The project is build using golang for webserver backend and sqlite3 for the database.

Folder structure is based on layers (db, domain, http) and every file will represent features (order.go, item.go, server.go, database.go, etc). Current project structure is:

```
+ db -> database source adapter
--+ database.go
--+ order.db
--+ order.go
+ domain -> the core of feature business
--+ item.go
--+ order.go
--+ result.go
+ http -> web source adapter
--+ server.go
--+ order.go
```

### Prerequisite

> Golang (~v1.21.3)
> gcc and g++ multi-lib
> any database client

### Run

server will be running at :9090 (change config at http/server.go:SERVER_ADDR)

> go run main.go

### Misc

you could try APIs using [Postman Collection](docs/Golang%20Beginner%20Class.postman_collection.json)