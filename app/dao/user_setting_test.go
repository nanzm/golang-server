package dao

import (
	"dora/pkg/utils"
	"testing"
)

func TestUserSettingDao_Get(t *testing.T) {
	dao := NewUserSettingDao()
	get, err := dao.Get(2)
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(get)
}
