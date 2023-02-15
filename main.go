package main

import (
	"SinaDownLoad/lib"
	"fmt"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Print("[+] Directory to be processed: [" + pwd + "]:")
	var path string
	fmt.Scanln(&path)
	if path == "" {
		path = pwd
	}

	var targetDir string
	fmt.Print("[+] Download location:[Download]:")
	fmt.Scanln(&targetDir)
	if targetDir == "" {
		targetDir = "Download"
	}

	files, err := lib.GetFileName(path)
	if err != nil {
		panic(err)
	}
	// get files

	var urlList []string
	for _, file := range files {
		urls, errFind := lib.UrlFinding(file)
		if errFind != nil {
			panic(errFind)
		}
		for _, url := range urls {
			urlList = append(urlList, url)
		}
	}
	// get Image URLs

	errDown := lib.BatchDownLoad(urlList, targetDir)
	// Here to change target dir

	if errDown != nil {
		panic(errDown)
	}

}
