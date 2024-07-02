package crud

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"gocrud/pkg/models"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func ReadData(db *gorm.DB, req *models.QueryRequest) (interface{}, error) {
	cacheKey := fmt.Sprintf("%v", req)
	if cachedResult, found := c.Get(cacheKey); found {
		return cachedResult, nil
	}

	resultType := reflect.TypeOf(req.Struct)
	if resultType.Kind() == reflect.Ptr {
		resultType = resultType.Elem()
	}
	resultSlice := reflect.MakeSlice(reflect.SliceOf(resultType), 0, 0)
	result := reflect.New(resultSlice.Type()).Interface()

	var conditionStrings []string
	var params []interface{}

	for key, value := range req.Conditions {
		if strings.Contains(key, " LIKE ") {
			conditionStrings = append(conditionStrings, fmt.Sprintf("%s ?", key))
		} else {
			conditionStrings = append(conditionStrings, fmt.Sprintf("%s = ?", key))
		}
		params = append(params, value)
	}

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

	if len(req.OrderBy) > 0 {
		query = fmt.Sprintf("%s ORDER BY %s", query, strings.Join(req.OrderBy, ", "))
	}

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
