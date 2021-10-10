package config

import (
	"fmt"
	"testing"
)

func init()  {
	MustLoad("../../.env")
}

func TestConfig(t *testing.T) {
	c := GetGorm()
	fmt.Printf("%#v \n", c)
}
