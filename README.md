# Spacecraft Management API

The Spacecraft Management API provides a set of RESTful API endpoints to create, retrieve, update, and delete spacecraft records. This project uses the Ent framework for schema definition and Go's native net/http package and gorilla/mux for routing and handler implementations.
For orm facebook opensource project ent is used here.

## Prerequisites

- Go 1.14 or newer
- MySQL 5.7 or newer (Docker recommended)

## Setting Up

1. Clone the repo
2. Setup db (If you're using docker then run below)

```
docker run --name=spacecraft -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=spacecraft -e MYSQL_USER=user -e MYSQL_PASSWORD=password mysql:latest
```

3. Run codegen

```
go generate ./ent
```

4. Make .env file

5. Run the server

```
go run ./cmd/spacecraft/main.go
```

## Testing

To execute all tests within the project, run:

```
go test ./...
```

## Available Routes

- GET /spaceships: Lists all spaceships
- GET /spaceships/{id}: Retrieves specific spaceship details
- POST /spaceships: Creates a new spaceship
- PUT /spaceships/{id}: Updates an existing spaceship
- DELETE /spaceships/{id}: Deletes an existing spaceship
