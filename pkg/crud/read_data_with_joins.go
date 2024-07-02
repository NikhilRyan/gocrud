package crud

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"gocrud/pkg/models"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func ReadDataWithJoins(db *gorm.DB, req *models.JoinRequest) (interface{}, error) {
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

	joinSelects := []string{}
	for _, join := range req.Joins {
		if len(join.Selects) > 0 {
			joinSelects = append(joinSelects, join.Selects...)
		}
	}
	if len(joinSelects) > 0 {
		columnsQuery = fmt.Sprintf("%s, %s", columnsQuery, strings.Join(joinSelects, ", "))
	}

	query := fmt.Sprintf("SELECT %s FROM %s", columnsQuery, req.Table)

	for _, join := range req.Joins {
		joinConditionStrings := []string{}
		for key, value := range join.Conditions {
			if strings.Contains(key, " LIKE ") {
				joinConditionStrings = append(joinConditionStrings, fmt.Sprintf("%s ?", key))
			} else {
				joinConditionStrings = append(joinConditionStrings, fmt.Sprintf("%s = ?", key))
			}
			params = append(params, value)
		}

		joinCondition := ""
		if len(joinConditionStrings) > 0 {
			joinCondition = fmt.Sprintf(" AND %s", strings.Join(joinConditionStrings, " AND "))
		}

		query = fmt.Sprintf("%s %s JOIN %s ON %s%s", query, join.JoinType, join.Table, join.OnCondition, joinCondition)
	}

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
