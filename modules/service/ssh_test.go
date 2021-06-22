package service

import (
	"dora/pkg/utils/logx"
	"testing"
)

func TestSshServerExec(t *testing.T) {
	conf := SshConfig{
		SshHost:     "121.41.82.251",
		SshUser:     "root",
		SshPassword: "",
		SshKey:      "",
		SshPort:     22,
		AuthType:    Password,
	}
	out, err := SshExecCommand(conf, "echo \"hello\"")
	if err != nil {
		panic(err)
	}
	logx.Printf("%s \n", out)
}
