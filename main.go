package main

import (
	"SinaDownLoad/lib"
)

func main() {
	path := "/Users/xiabee/Desktop/GitHub/gitpage/source/_posts"
	// change your path here

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

	errDown := lib.BatchDownLoad(urlList, "Download")
	// Here to change target dir

	if errDown != nil {
		panic(errDown)
	}

}
