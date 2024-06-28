package crud

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type QueryRequest struct {
	Table      string                 `json:"table" binding:"required"`
	Columns    []string               `json:"columns"`
	Conditions map[string]interface{} `json:"conditions"`
	OrderBy    []string               `json:"order_by"`
	Limit      int                    `json:"limit"`
	Offset     int                    `json:"offset"`
	Struct     interface{}            `json:"struct"`
}

func ReadData(db *gorm.DB, req *QueryRequest) (interface{}, error) {
	cacheKey := fmt.Sprintf("%v", req)
	if cachedResult, found := c.Get(cacheKey); found {
		return cachedResult, nil
	}

	result := reflect.New(reflect.TypeOf(req.Struct)).Interface()
	var conditionStrings []string
	var params []interface{}

	for key, value := range req.Conditions {
		conditionStrings = append(conditionStrings, fmt.Sprintf("%s = ?", key))
		params = append(params, value)
	}

	// Infer columns from struct if provided
	columnsQuery := "*"
	if req.Struct != nil {
		columnsQuery = getStructFields(req.Struct)
	} else if len(req.Columns) > 0 {
		columnsQuery = strings.Join(req.Columns, ", ")
	}

	query := fmt.Sprintf("SELECT %s FROM %s", columnsQuery, req.Table)
	if len(conditionStrings) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(conditionStrings, " AND "))
	}

	// Order by
	if len(req.OrderBy) > 0 {
		query = fmt.Sprintf("%s ORDER BY %s", query, strings.Join(req.OrderBy, ", "))
	}

	// Limit and offset
	if req.Limit > 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, req.Limit)
	}

	if req.Offset > 0 {
		query = fmt.Sprintf("%s OFFSET %d", query, req.Offset)
	}

	if err := db.Raw(query, params...).Scan(result).Error; err != nil {
		return nil, err
	}

	c.Set(cacheKey, result, cache.DefaultExpiration)

	return result, nil
}

func getStructFields(obj interface{}) string {
	t := reflect.TypeOf(obj)
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("gorm")
		if tag != "" && strings.HasPrefix(tag, "column:") {
			columns := strings.Split(tag, ":")
			fields = append(fields, columns[1])
		} else {
			fields = append(fields, field.Name)
		}
	}
	return strings.Join(fields, ", ")
}
