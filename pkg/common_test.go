package pkg

import (
	"fmt"
	"testing"
)

type newObj struct {
	Id   float64 `json:"id"`
	Name string  `json:"name"`
	Age  int     `json:"age"`
}

func TestWeekDecode(t *testing.T) {
	f := map[string]string{
		"id":   "null",
		"age":  "null",
		"name": "null",
	}

	obj := &newObj{}
	err := WeekDecode(f, obj)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v \n", obj)
}
