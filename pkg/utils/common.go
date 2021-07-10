package utils

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"strconv"
)

func Paginate(current, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		c := 1
		if current > 0 {
			c = int(current)
		}

		s := 10
		if pageSize > 0 {
			s = int(pageSize)
		}

		return db.Limit(s).Offset((c - 1) * s).Order("id desc")
	}
}

func WeekDecode(input interface{}, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &output,
		TagName:          "json",
		DecodeHook:       mapstructure.ComposeDecodeHookFunc(stringToFloatCheckHookFunc()),
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(input)
	if err != nil {
		return err
	}
	return nil
}

func stringToFloatCheckHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		switch f.Kind() {
		case reflect.String:
			if data.(string) == "null" && t.Kind() != reflect.String {
				return 0, nil
			}
			return data, nil
		default:
			return data, nil
		}
	}
}

func MatchStackLineCol(stackLine string) (line, col int, e error) {
	r := regexp.MustCompile(".+:(\\d+):(\\d+).+")
	sub := r.FindStringSubmatch(stackLine)
	if len(sub) < 3 {
		return 0, 0, errors.New("无法从堆栈中获得行列号")
	}

	var err error
	l, err := strconv.Atoi(sub[1])
	c, err := strconv.Atoi(sub[2])
	if err != nil {
		return 0, 0, err
	}
	return l, c, nil
}
