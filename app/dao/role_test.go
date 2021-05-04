package dao

import (
	"dora/pkg/logger"
	"dora/pkg/utils"
	"testing"
)

func TestRoleDao_Create(t *testing.T) {

}

func TestRoleDao_Delete(t *testing.T) {

}

func TestRoleDao_Get(t *testing.T) {
	dao := NewRoleDao()
	get, err := dao.Get(1)
	if err != nil {
		panic(err)
	}

	utils.PrettyPrint(get)
}

func TestRoleDao_List(t *testing.T) {
	dao := NewRoleDao()
	list, current, size, total, err := dao.List(1, 3)
	if err != nil {
		t.Fatal(err)
	}
	logger.Printf("current: %v size: %v total %v \n", current, size, total)
	utils.PrettyPrint(list)
}

func TestRoleDao_Update(t *testing.T) {

}
