package models

type QueryRequest struct {
	Table      string                 `json:"table" binding:"required"`
	Columns    []string               `json:"columns"`
	Conditions map[string]interface{} `json:"conditions"`
	OrderBy    []string               `json:"order_by"`
	Limit      int                    `json:"limit"`
	Offset     int                    `json:"offset"`
	Struct     interface{}            `json:"struct"`
}

type JoinClause struct {
	JoinType    string                 `json:"join_type"`
	Table       string                 `json:"table"`
	OnCondition string                 `json:"on_condition"`
	Selects     []string               `json:"selects"`
	Conditions  map[string]interface{} `json:"conditions"`
}

type JoinRequest struct {
	Table      string                 `json:"table" binding:"required"`
	Joins      []JoinClause           `json:"joins"`
	Columns    []string               `json:"columns"`
	Conditions map[string]interface{} `json:"conditions"`
	OrderBy    []string               `json:"order_by"`
	Limit      int                    `json:"limit"`
	Offset     int                    `json:"offset"`
	Struct     interface{}            `json:"struct"`
}
