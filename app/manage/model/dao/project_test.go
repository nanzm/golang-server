package dao

import (
	"dora/app/manage/model/entity"
	"dora/pkg/utils"
	"dora/pkg/utils/logx"
	"testing"
)

func TestProjectDao_Create(t *testing.T) {
	dao := NewProjectDao()

	p := entity.Project{
		AppId:             utils.RandString(10),
		Name:              utils.RandString(5),
		Alias:             "3",
		Type:              "4",
		GitRepositoryUrl:  "5",
		GitRepositoryName: "6",
	}
	create, err := dao.Create(&p, 2)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(create)
}

func TestProjectDao_Update(t *testing.T) {
	dao := NewProjectDao()
	newVal := entity.Project{
		ID:                9,
		Name:              utils.RandString(15),
		Alias:             "",
		Type:              "",
		GitRepositoryUrl:  "",
		GitRepositoryName: "",
	}
	err := dao.Update(newVal)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(newVal)
}

func TestProjectDao_Delete(t *testing.T) {
	dao := NewProjectDao()
	err := dao.Delete(6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProjectDao_Get(t *testing.T) {
	dao := NewProjectDao()
	get, err := dao.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(get)
}

func TestProjectDao_List(t *testing.T) {
	dao := NewProjectDao()
	list, current, size, total, err := dao.List(1, 10)

	if err != nil {
		panic(err)
	}

	logx.Printf("current: %v size: %v total %v \n", current, size, total)
	utils.PrettyPrint(list)
}

func TestProjectDao_ProjectUsers(t *testing.T) {
	dao := NewProjectDao()
	users, err := dao.ProjectUsers(8)
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(users)
}
