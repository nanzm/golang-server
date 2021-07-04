package elastic

import (
	"dora/pkg/utils"
	"strings"
	"testing"
)

func TestGetElasticClient(t *testing.T) {
	es := GetClient()
	info, err := es.Info()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestPut(t *testing.T) {
	es := GetClient()

	cp := ``
	mapping, err2 := es.Indices.Create("dora_test", es.Indices.Create.WithBody(strings.NewReader(cp)))

	if err2 != nil {
		t.Fatal(err2)
	}

	utils.PrettyPrint(mapping)
}
