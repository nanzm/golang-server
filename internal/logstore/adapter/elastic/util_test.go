package elasticComponent

import (
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

func TestGetFirstAndLastTime(t *testing.T) {
	data := `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 12312312, "nets": ["ig", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 41237, "nets": ["ig", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 427, "nets": ["ig", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 4312353247, "nets": ["ig", "tw"]}
  ]
}`

	list := gjson.Get(data, "friends.#.age").Array()
	fmt.Printf("%v \n", list)

	f, l := GetFirstAndLastTime(list)
	fmt.Printf("%v %v\n", f, l)
}
