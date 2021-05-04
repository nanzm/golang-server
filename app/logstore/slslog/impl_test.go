package slslogComponent

import (
	"dora/config"
	"dora/pkg/utils"
	"testing"
)

func init() {
	config.ParseConf("../../../config.yml")
}
func Test_slsQuery_GetLogByMd5(t *testing.T) {
	log, err := NewSlsQuery().GetLogByMd5(1617120000, 1617206399,
		"36dce4bc005e0e12b2c66e9ac7cbbbe4")

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_LogCountByMd5(t *testing.T) {
	log, err := NewSlsQuery().LogCountByMd5(1614652538, 1617330951,
		"a2efea7daa4c9bd325d8b28617d06a6c")

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_PvUvTrend(t *testing.T) {
	log, err := NewSlsQuery().PvUvTrend("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951, 600)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_SdkVersionCount(t *testing.T) {
	log, err := NewSlsQuery().SdkVersionCount("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_PagesCount(t *testing.T) {
	log, err := NewSlsQuery().PagesCount("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ErrorCount(t *testing.T) {
	log, err := NewSlsQuery().ErrorCount("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ErrorCountTrend(t *testing.T) {
	log, err := NewSlsQuery().ErrorCountTrend("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951, 1200)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ApiErrorCount(t *testing.T) {
	log, err := NewSlsQuery().ApiErrorCount("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ApiErrorTrend(t *testing.T) {
	log, err := NewSlsQuery().ApiErrorTrend("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951, 1200)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ApiErrorList(t *testing.T) {
	log, err := NewSlsQuery().ApiErrorList("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_PerfNavigationTimingTrend(t *testing.T) {
	log, err := NewSlsQuery().PerfNavigationTimingTrend("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1617292800, 1617373353, 720)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_PerfNavigationTimingValues(t *testing.T) {
	log, err := NewSlsQuery().PerfNavigationTimingValues("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ProjectEnv(t *testing.T) {
	log, err := NewSlsQuery().ProjectEnv("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ProjectUserScreen(t *testing.T) {
	log, err := NewSlsQuery().ProjectUserScreen("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ProjectVersion(t *testing.T) {
	log, err := NewSlsQuery().ProjectVersion("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ProjectSendMode(t *testing.T) {
	log, err := NewSlsQuery().ProjectSendMode("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ProjectEventCount(t *testing.T) {
	log, err := NewSlsQuery().ProjectEventCount("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_ResLoadFailList(t *testing.T) {
	log, err := NewSlsQuery().ResLoadFailList("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}

func Test_slsQuery_CategoryCount(t *testing.T) {
	log, err := NewSlsQuery().CategoryCount("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1614652538, 1617330951)

	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(log)
}