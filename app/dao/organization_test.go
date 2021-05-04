package dao

import (
	"dora/app/model"
	"dora/pkg/logger"
	"dora/pkg/utils"
	"testing"
)

func TestOrganizationDao_Create(t *testing.T) {
	dao := NewOrganizationDao()

	p := model.Organization{
		Name:         utils.RandString(10),
		Introduction: utils.RandString(100),
		Type:         "",
		CreateUid:    1,
	}
	create, err := dao.Create(1, &p)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(create)
}

func TestOrganizationDao_Update(t *testing.T) {

}

func TestOrganizationDao_Delete(t *testing.T) {

}

func TestOrganizationDao_Get(t *testing.T) {
	dao := NewOrganizationDao()
	get, err := dao.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(get)
}

func TestOrganizationDao_List(t *testing.T) {
	dao := NewOrganizationDao()
	list, current, size, total, err := dao.List(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	logger.Printf("current: %v size: %v total %v \n", current, size, total)
	utils.PrettyPrint(list)
}

func TestOrganizationDao_AddUser(t *testing.T) {
	//dao := NewOrganizationDao()
	//err := dao.AddUser(19, 13)
	//if err != nil {
	//	t.Fatal(err)
	//}
}

func TestOrganizationDao_RemoveUser(t *testing.T) {
	dao := NewOrganizationDao()
	err := dao.RemoveUser(19, 13)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrganizationDao_UserOrganization(t *testing.T) {
	dao := NewOrganizationDao()
	organizations, err := dao.UserOrganization(1)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(organizations)
}

func TestOrganizationDao_GetMembers(t *testing.T) {
	dao := NewOrganizationDao()
	organizations, err := dao.GetMembers(1)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(organizations)
}
