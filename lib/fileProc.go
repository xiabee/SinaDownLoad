package lib

import (
	"os"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// to check if the path exists

func DirCheck(dir string) error {
	exist, err := PathExists(dir)
	if !exist {
		errdir := os.Mkdir(dir, os.ModePerm)
		if errdir != nil {
			return errdir
		}
	}
	return err
}

// to check and create the Dir

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// to check the dir

func PathProc(path string) (string, error) {
	_, err := PathExists(path)
	if path[len(path)-1:] != "/" {
		path += "/"
	}
	return path, err
}

func GetFileName(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !IsDir(path) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// get the files in certain dir
