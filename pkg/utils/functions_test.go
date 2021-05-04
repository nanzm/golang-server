package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandom(t *testing.T) {

	for i := 10; i < 1000000; i++ {
		toString, err := EncodeToString(6)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%v \n", toString)
	}

}

func TestUid(t *testing.T) {
	fmt.Printf("%v \n", uuid.New())
	fmt.Printf("%v \n", uuid.New())
	fmt.Printf("%v \n", uuid.New())
}

func TestRemoveDuplicateElement(t *testing.T) {
	a := []string{
		"c0c74bcd080c3a2602e41c1013862ae7",
		"c0c74bcd080c3a2602e41c1013862ae7",
		"c0c74bcd080c3a2602e41c1013862ae7",
		"c0c74bcd080c3a2602e41c1013862ae7",
		"c0c74bcd080c3a2602e41c1013862ae7",
		"c0c74bcd080c3a2602e41c1013862ae7"}

	element := Uniq(a)
	fmt.Printf("%v \n", element)
}

func TestMd5(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"3", args{[]byte("")}, "d41d8cd98f00b204e9800998ecf8427e"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.data); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeJsonMarshal(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"空字符串", args{""}, ""},
		{"nil", args{nil}, ""},
		{"结构体", args{struct{ M string }{"haha"}}, "{\"M\":\"haha\"}"},
		{"map", args{map[string]string{"1": "2"}}, "{\"1\":\"2\"}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeJsonMarshal(tt.args.in); got != tt.want {
				t.Errorf("SafeJsonMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceKeyPrefix(t *testing.T) {
	testData := map[string]interface{}{
		"a":    "12",
		"_b":   "123",
		"_c_z": "1234",
	}

	result := ReplaceKeyPrefix(testData, "_", "d_")
	assert.Equal(t, result, map[string]interface{}{
		"a":     "12",
		"d_b":   "123",
		"d_c_z": "1234",
	}, "they should be equal")

	result2 := ReplaceKeyPrefix(testData, "_", "")
	assert.Equal(t, result2, map[string]interface{}{
		"a":   "12",
		"b":   "123",
		"c_z": "1234",
	}, "they should be equal")
}
