# GoCRUD

## Overview
This library provides a dynamic CRUD interface that allows users to perform operations without writing any code, just by providing the table name and key.

## Features
- Dynamic table creation
- Insert, read, update, and delete operations
- Configurable via simple JSON requests
- RESTful API built with Gin and GORM

## Getting Started

### Prerequisites
- Go 1.17 or later
- Docker (optional, for containerization)

### Installation
Clone the repository:
```sh
git clone https://github.com/nikhilryan5/gocrud.git
cd gocrud
```

## Running the application

### Using Go
```sh
go mod download
go build -o main ./cmd/main.go
./main
```

### Using Docker
```sh
docker build -t gocrud .
docker run -p 8080:8080 gocrud
```

### Usage
Send JSON requests to the endpoints:
- POST `/create`
- GET `/read`
- PUT `/update`
- DELETE `/delete`

### Example
#### Create
```sh
curl -X POST "http://localhost:8080/create" -H "Content-Type: application/json" -d '{
    "table": "users",
    "key": "id",
    "data": {
        "name": "John Doe",
        "age": 30
    }
}'
```

#### Read
```sh
curl -X GET "http://localhost:8080/read" -H "Content-Type: application/json" -d '{
    "table": "users",
    "key": "id",
    "data": {
        "id": 1
    }
}'
```

#### Update
```sh
curl -X PUT "http://localhost:8080/update" -H "Content-Type: application/json" -d '{
    "table": "users",
    "key": "id",
    "data": {
        "id": 1,
        "name": "Jane Doe",
        "age": 25
    }
}'
```

#### Delete
```sh
curl -X DELETE "http://localhost:8080/delete" -H "Content-Type: application/json" -d '{
    "table": "users",
    "key": "id",
    "data": {
        "id": 1
    }
}'
```

## License

This project is licensed under the MIT License.