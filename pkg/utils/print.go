package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyJson(b []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", out.Bytes())
}

func PrettyString(in interface{}) string {
	b, err := json.Marshal(in)
	if err != nil {
		return err.Error()
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		return err.Error()
	}
	return out.String()
}

func PrettyPrint(in interface{}) {
	b, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v \n", out.String())
	//
	//c := color.New(color.Bold, color.FgGreen).Add(color.BgWhite)
	//c.EnableColor()
	//_, err = c.Println(out.String())
	//if err != nil {
	//	panic(err)
	//}
}

func ColorPrint() {
	//rb := new(bytes.Buffer)
}
