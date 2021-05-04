package elasticComponent

import (
	"dora/pkg/utils"
	"testing"
)

func Test_elasticQuery_PvUvTotal(t *testing.T) {
	total, err := NewElasticQuery().PvUvTotal("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1617094800, 1617102249)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(total)
}

func Test_elasticQuery_PvUvTrend(t *testing.T) {
	re, err := NewElasticQuery().PvUvTrend("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1617097374, 1617183785, 30)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(re)
}

func Test_elasticQuery_EntryPage(t *testing.T) {
	re, err := NewElasticQuery().PagesCount("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1617097374, 1617183785)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(re)
}

//
//func Test_elasticQuery_ErrorCount(t *testing.T) {
//	re, err := NewElasticQuery().ErrorCount("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1617097374, 1617183785)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	utils.PrettyPrint(re)
//}
