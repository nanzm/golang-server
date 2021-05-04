package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetStackSourceMap(dir string, stack string) (sourcemap []byte, err error) {
	var f string
	e := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ext := path.Ext(info.Name())
		// 只检查 .map 文件
		if ext == ".map" {
			// 去掉后缀 再匹配
			filePrefix := strings.TrimSuffix(info.Name(), ext)
			index := strings.Index(stack, filePrefix)
			if index > -1 {
				f = p
			}
		}
		return nil
	})

	if e != nil {
		return nil, e
	}
	if f == "" {
		return nil, errors.New("未找到对应 sourcemap 文件")
	}

	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, errors.New("读取 sourcemap 文件失败")
	}
	return file, nil
}
