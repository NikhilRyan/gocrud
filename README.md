
### Usage
Send JSON requests to the endpoints:
- POST `/create`
- GET `/read`
- POST `/read-with-joins`
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

#### Read with Conditions, Order By, Limit, Offset, and Struct (with Caching)
```sh
curl -X GET "http://localhost:8080/read" -H "Content-Type: application/json" -d '{
    "table": "users",
    "conditions": {
        "age": 30
    },
    "order_by": ["name ASC", "age DESC"],
    "limit": 10,
    "offset": 5,
    "struct": {"ID": 0, "Name": "", "Age": 0}
}'
```

#### Read with Joins, Conditions, Order By, Limit, Offset, and Struct (with Caching)
```sh
curl -X POST "http://localhost:8080/read-with-joins" -H "Content-Type: application/json" -d '{
    "table": "orders",
    "joins": [
        {
            "join_type": "INNER",
            "table": "users",
            "on_condition": "orders.user_id = users.id",
            "selects": ["users.name", "users.email"],
            "conditions": {
                "users.age": 30
            }
        },
        {
            "join_type": "LEFT",
            "table": "products",
            "on_condition": "orders.product_id = products.id",
            "selects": ["products.name as product_name"],
            "conditions": {
                "products.price >": 100
            }
        }
    ],
    "columns": ["orders.id", "orders.total"],
    "conditions": {
        "orders.status": "completed",
        "orders.created_at >": "2022-01-01",
        "orders.note LIKE": "%urgent%"
    },
    "order_by": ["orders.date DESC"],
    "limit": 10,
    "offset": 5,
    "struct": {"OrderID": 0, "Total": 0, "Name": "", "Email": "", "ProductName": ""}
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

#### Get Column Info
```sh
curl -X GET "http://localhost:8080/columns?table=users"
```

### Running Tests
Unit tests are provided to verify the functionality of each handler. The test files are located in the `pkg/handlers` directory.

#### Run All Tests
```sh
go test ./...
```

#### Test Cases for ReadWithJoinsHandler
```go
package handlers

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "bytes"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gocrud/pkg/crud"
)

type OrderWithUser struct {
    OrderID     int    `gorm:"column:order_id"`
    Total       int    `gorm:"column:total"`
    Name        string `gorm:"column:name"`
    Email       string `gorm:"column:email"`
    ProductName string `gorm:"column:product_name"`
}

func TestReadWithJoinsHandler(t *testing.T) {
    // Setup
    db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    
    router := gin.Default()
    router.POST("/read-with-joins", ReadWithJoinsHandler(db))

    // Test case
    reqBody := `{
        "table": "orders",
        "joins": [
            {
                "join_type": "INNER",
                "table": "users",
                "on_condition": "orders.user_id = users.id",
                "selects": ["users.name", "users.email"],
                "conditions": {
                    "users.age": 30
                }
            },
            {
                "join_type": "LEFT",
                "table": "products",
                "on_condition": "orders.product_id = products.id",
                "selects": ["products.name as product_name"],
                "conditions": {
                    "products.price >": 100
                }
            }
        ],
        "columns": ["orders.id", "orders.total"],
        "conditions": {
            "orders.status": "completed",
            "orders.created_at >": "2022-01-01",
            "orders.note LIKE": "%urgent%"
        },
        "order_by": ["orders.date DESC"],
        "limit": 10,
        "offset": 5,
        "struct": OrderWithUser{}
    }`
    req, _ := http.NewRequest("POST", "/read-with-joins", bytes.NewBufferString(reqBody))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}
```
