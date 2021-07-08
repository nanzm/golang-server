package fs

import (
	"fmt"
	"os"
)

func FileExists(filePath string) (bool, error) {
	st, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return false, err
		}
		return false, nil
	}

	if st.IsDir() {
		return false, fmt.Errorf("%s is a directory", filePath)
	}

	return true, nil
}

func EnsureDir(dirPath string) error {
	var err error
	if _, err = os.Stat(dirPath); err != nil && os.IsNotExist(err) {
		e := os.MkdirAll(dirPath, os.ModePerm)
		if e != nil {
			return e
		}
		return nil
	}
	return err
}

func FileSize(filePath string) int64 {
	s, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return s.Size()
}
