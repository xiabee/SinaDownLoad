package lib

import (
	"fmt"
	"os"
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

func GetFileName(path string) ([]os.DirEntry, error) {
	path, _ = PathProc(path)
	files, err := os.ReadDir(path)
	var result []os.DirEntry
	for _, f := range files {
		if IsDir(path + f.Name() + "/") {
			GetFileName(path + f.Name() + "/")
		} else {
			fmt.Println(path + f.Name())
			result = append(result, f)
		}
	}
	return result, err
}
