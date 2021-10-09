package elasticComponent

import (
	"dora/pkg/utils"
	"dora/pkg/utils/logx"
	"testing"
)

func init() {
	logx.Init("./dora.log")
}

func Test_elasticQuery(t *testing.T) {
	re, err := NewElasticQuery().GetErrorLogsByMd5("wdssfar2312312dsad", 1624177745000, 1625041745000, "074aa2b51cb6b17319bc52b9a3a18d45")
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(re)
}
