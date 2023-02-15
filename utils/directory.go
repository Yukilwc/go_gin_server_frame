package utils

import (
	"errors"
	"os"
)

// 文件目录是否存在

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		if fi.IsDir() {
			return true, nil
		} else {
			return false, errors.New("存在同名文件")
		}
	}
}
