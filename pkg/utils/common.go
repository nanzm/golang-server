package utils

import (
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"reflect"
)

func Paginate(current, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		c := 1
		if current > 0 {
			c = current
		}

		s := 10
		if pageSize > 0 {
			s = pageSize
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
