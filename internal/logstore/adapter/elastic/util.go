package elasticComponent

import (
	"github.com/tidwall/gjson"
	"sort"
)

func GetFirstAndLastTime(arr []gjson.Result) (first int, last int) {
	list := make([]int, 0)
	for _, item := range arr {
		list = append(list, int(item.Num))
	}
	sort.Ints(list)

	if len(list) == 0 {
		return 0, 0
	}

	if len(list) == 1 {
		return list[0], list[0]
	}

	return list[0], list[len(list)-1]
}
