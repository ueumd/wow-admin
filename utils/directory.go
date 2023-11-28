package utils

import (
	"errors"
	"os"
)

func PathExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if fileInfo.IsDir() {
		return true, nil
	}
	return false, errors.New("already exists")
}
