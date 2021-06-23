package mqConsumer

import (
	"reflect"
	"testing"
)

func Test_md5AggData(t *testing.T) {
	type args struct {
		event map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"没有agg", args{map[string]interface{}{}}, ""},
		{"agg nil", args{map[string]interface{}{"agg": nil}}, ""},
		{"agg 空字符串", args{map[string]interface{}{"agg": ""}}, ""},
		{"agg 有内容", args{map[string]interface{}{"agg": "咚咚咚咚咚"}}, "600aa0ab0200d06400c3eadc60dc7b26"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5AggData(tt.args.event); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genEventMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}
