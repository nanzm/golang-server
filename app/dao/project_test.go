package dao

import (
	"dora/app/model"
	"dora/config"
	"dora/pkg/logger"
	"dora/pkg/utils"
	"testing"
)

func init() {
	config.ParseConf("../../config.yml")
}

func TestProjectDao_Create(t *testing.T) {
	dao := NewProjectDao()

	p := model.Project{
		AppId:             utils.RandString(10),
		Name:              utils.RandString(5),
		Alias:             "3",
		Type:              "4",
		GitRepositoryUrl:  "5",
		GitRepositoryName: "6",
		OrganizationId:    1,
	}
	create, err := dao.Create(&p)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(create)
}

func TestProjectDao_Update(t *testing.T) {
	dao := NewProjectDao()
	newVal := model.Project{
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

	logger.Printf("current: %v size: %v total %v \n", current, size, total)
	utils.PrettyPrint(list)
}

func TestProjectDao_OrganizationProjectsList(t *testing.T) {
	dao := NewProjectDao()
	list, err := dao.OrganizationProjectsList(1)
	if err != nil {
		panic(err)
	}

	utils.PrettyPrint(list)
}
