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
git clone https://github.com/nikhilryan/gocrud.git
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

### LICENSE

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
