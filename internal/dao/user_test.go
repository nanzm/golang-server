package dao

import (
	"dora/internal/model"
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
		RoleId:   1,
	}
	err := dao.Update(&user)
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(user)
}

func TestUserDao_UserProjects(t *testing.T) {
	dao := NewUserDao()
	data, err := dao.UserProjects(1)
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(data)
}

func TestUserDao_GetByName(t *testing.T) {
	dao := NewUserDao()
	user, err := dao.GetByName("dora")
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(user)
}

func TestUserDao_UserProjects1(t *testing.T) {
	dao := NewUserDao()
	user, err := dao.UserProjects(1)
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(user)
}