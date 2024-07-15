package crud

import (
	"fmt"
	"gocrud/pkg/cache"
	"gocrud/pkg/models"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func ReadDataWithJoins(db *gorm.DB, req *models.JoinRequest) (interface{}, error) {
	cacheKey := fmt.Sprintf("read_with_joins:%v", req)
	result, err := cache.ReadFromCache(cacheKey, req.Struct, func() (interface{}, error) {
		query, params := GenerateReadWithJoinsQuery(req)
		result := reflect.New(reflect.SliceOf(reflect.TypeOf(req.Struct).Elem())).Interface()
		if err := db.Raw(query, params...).Scan(result).Error; err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GenerateReadWithJoinsQuery(req *models.JoinRequest) (string, []interface{}) {
	var conditionStrings []string
	var params []interface{}

	// Process main table conditions
	for key, value := range req.Conditions {
		if strings.Contains(key, " LIKE ") {
			conditionStrings = append(conditionStrings, fmt.Sprintf("%s ?", key))
		} else {
			conditionStrings = append(conditionStrings, fmt.Sprintf("%s = ?", key))
		}
		params = append(params, value)
	}

	columnsQuery := strings.Join(req.Columns, ", ")
	if req.Struct != nil && len(req.Columns) == 0 {
		columnsQuery = getStructFields(req.Struct)
	}

	var joinSelects []string
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
		var joinConditionStrings []string
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

	return query, params
}
