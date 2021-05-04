package dao

import (
	"dora/app/model"
	"dora/pkg/logger"
	"dora/pkg/utils"
	"testing"
)

func TestUserDao_Create(t *testing.T) {
	dao := NewUserDao()
	user := model.User{
		Password: utils.RandString(12),
		NickName: utils.RandString(12),
		Avatar:   utils.RandString(12),
		Email:    utils.RandString(12),
		Phone:    utils.RandString(12),
		Status:   0,
		RoleId:   2,
	}

	create, err := dao.Create(&user)
	if err != nil {
		panic(err)
	}

	utils.PrettyPrint(create)
}

func TestUserDao_Delete(t *testing.T) {

}

func TestUserDao_Get(t *testing.T) {
	dao := NewUserDao()
	get, err := dao.Get(3)
	if err != nil {
		panic(err)
	}

	utils.PrettyPrint(get)
}

func TestUserDao_List(t *testing.T) {
	dao := NewUserDao()
	list, current, size, total, err := dao.List(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	logger.Printf("current: %v size: %v total %v \n", current, size, total)
	utils.PrettyPrint(list)
}

func TestUserDao_Update(t *testing.T) {
	dao := NewUserDao()

	user := model.User{
		ID:       2,
		NickName: "4",
		Password: "3",
		Avatar:   "5",
		Email:    "6",
		Phone:    "",
		RoleId:   1,
	}
	err := dao.Update(&user)
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(user)
}

func TestProjectDao_UserOrganizations(t *testing.T) {
	dao := NewUserDao()
	organizations, err := dao.UserOrganizations(1)
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(organizations)
}

func TestUserDao_GetByName(t *testing.T) {
	dao := NewUserDao()
	user, err := dao.GetByName("dora")
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(user)
}
