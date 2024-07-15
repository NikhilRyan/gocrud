package crud

import (
	"gocrud/pkg/models"
	"reflect"
	"testing"
)

func TestGenerateReadQuery(t *testing.T) {
	req := &models.QueryRequest{
		Table:      "users",
		Columns:    []string{"id", "name", "age"},
		Conditions: map[string]interface{}{"age >": 30},
		OrderBy:    []string{"name ASC"},
		Limit:      10,
		Offset:     0,
		Struct:     &[]models.User{},
	}
	query, params := GenerateReadQuery(req)
	expectedQuery := "SELECT id, name, age, email FROM users WHERE age > = ? ORDER BY name ASC LIMIT 10"
	if query != expectedQuery {
		t.Errorf("expected query: %s, got: %s", expectedQuery, query)
	}
	expectedParams := []interface{}{30}
	if !reflect.DeepEqual(params, expectedParams) {
		t.Errorf("expected params: %v, got: %v", expectedParams, params)
	}
}

func TestGenerateReadWithJoinsQuery(t *testing.T) {
	req := &models.JoinRequest{
		Table: "orders",
		Joins: []models.JoinClause{
			{
				JoinType:    "INNER",
				Table:       "users",
				OnCondition: "orders.user_id = users.id",
				Selects:     []string{"users.name", "users.email"},
				Conditions:  map[string]interface{}{"users.age >": 30},
			},
			{
				JoinType:    "LEFT",
				Table:       "products",
				OnCondition: "orders.product_id = products.id",
				Selects:     []string{"products.name AS product_name"},
				Conditions:  map[string]interface{}{"products.price >": 100},
			},
		},
		Columns:    []string{"orders.id", "orders.total"},
		Conditions: map[string]interface{}{"orders.status": "completed"},
		OrderBy:    []string{"orders.date DESC"},
		Limit:      10,
		Offset:     5,
		Struct:     &[]models.Order{},
	}
	query, params := GenerateReadWithJoinsQuery(req)
	expectedQuery := "SELECT orders.id, orders.total, users.name, users.email, products.name AS product_name FROM orders INNER JOIN users ON orders.user_id = users.id AND users.age > = ? LEFT JOIN products ON orders.product_id = products.id AND products.price > = ? WHERE orders.status = ? ORDER BY orders.date DESC LIMIT 10 OFFSET 5"
	if query != expectedQuery {
		t.Errorf("expected query: %s, got: %s", expectedQuery, query)
	}
	expectedParams := []interface{}{"completed", 30, 100}
	if !reflect.DeepEqual(params, expectedParams) {
		t.Errorf("expected params: %v, got: %v", expectedParams, params)
	}
}
