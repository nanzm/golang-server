package service

import (
	"dora/pkg/utils/logx"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"time"
)

const (
	Password string = "password"
	Key      string = "key"
)

type SshConfig struct {
	SshHost     string
	SshUser     string
	SshPassword string
	SshKey      string
	SshPort     int
	AuthType    string
}

func SshExecCommand(conf SshConfig, cmd string) (output []byte, err error) {
	config := &ssh.ClientConfig{
		Timeout:         3 * time.Second,
		User:            conf.SshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	if conf.AuthType == Password {
		config.Auth = []ssh.AuthMethod{ssh.Password(conf.SshPassword)}
	}

	if conf.AuthType == Key {
		config.Auth = []ssh.AuthMethod{getKeys([]byte(conf.SshKey))}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", conf.SshHost, conf.SshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, errors.Wrap(err, "创建ssh client 失败")
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		return nil, errors.Wrap(err, "创建ssh session 失败")
	}
	defer session.Close()

	//执行远程命令
	combo, err := session.CombinedOutput(cmd)
	//combo, err := session.CombinedOutput("whoami; ls -al;echo https://github.com/dejavuzhou/felix")
	if err != nil {
		return nil, errors.Wrap(err, "远程执行cmd 失败")
	}
	return combo, nil
}

func getKeys(key []byte) ssh.AuthMethod {
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		logx.Panicf("ssh key signer failed %v", err)
	}
	return ssh.PublicKeys(signer)
}
