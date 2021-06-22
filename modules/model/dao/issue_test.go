package dao

import (
	"dora/pkg/utils"
	"testing"
)

func TestIssueDao_Create(t *testing.T) {

}

func TestIssueDao_Delete(t *testing.T) {

}

func TestIssueDao_Get(t *testing.T) {

}

func TestIssueDao_List(t *testing.T) {
	//dao := NewIssueDao()
	//list, current, size, total, err := dao.ListQueryTimeRange("", 1, 10)
	//if err != nil {
	//	panic(err)
	//}
	//logx.Printf("current: %v size: %v total %v \n", current, size, total)
	//utils.PrettyPrint(list)
}

func TestIssueDao_Update(t *testing.T) {

}

func TestIssueDao_Query(t *testing.T) {
	dao := NewIssueDao()
	query, err := dao.QueryByMd5("ddd")
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(query)
}

func TestIssueDao_ListQueryTimeRange(t *testing.T) {
	//dao := NewIssueDao()
	//gotResult, gotCurrent, gotPageSize, gotTotal, err := dao.ListQueryTimeRange("fca5deec-a9db-4dac-a4db-b0f4610d16a5",
	//	"2020-02-01 00:00:00", "2020-02-01 23:59:59", 1, 10)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//utils.PrettyPrint(gotResult)
	//utils.PrettyPrint(gotCurrent)
	//utils.PrettyPrint(gotPageSize)
	//utils.PrettyPrint(gotTotal)
}
