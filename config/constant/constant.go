package constant

const (
	AggregatesAvg = "avg"
	AggregatesSum = "sum"
	AggregatesMin = "min"
	AggregatesMax = "max"

	OperatorGt  = "gt"  // > 大于
	OperatorLt  = "lt"  // < 小于
	OperatorEq  = "eq"  // == 等于
	OperatorGte = "gte" // >= 大于等于
	OperatorLte = "lte" // <= 小于等于

	MeasureError = "error"
	MeasureApi   = "api"
	MeasureRes   = "res"
)

const (
	Fail    = -1
	Success = 0
	BizMsg  = 3000
)
