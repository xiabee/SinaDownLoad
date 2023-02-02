package main

import (
	"SinaDownLoad/lib"
	"fmt"
)

func main() {
	path := "/Users/xiabee/Desktop/GitHub/gitpage/source/_posts"

	files, err := lib.GetFileName(path)
	if err != nil {
		panic(err)
	}

	var urlList []string
	for _, file := range files {
		urls, _ := lib.UrlFinding(file)
		for _, url := range urls {
			urlList = append(urlList, url)
		}
		for _, url := range urlList {
			fmt.Println(url, file)
		}
	}

	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, url := range urlList {
	//	fmt.Println(url)
	//}
	//url := "https://tva1.sinaimg.cn/large/0084b03xgy1h58pm1gs0uj30y20iu0za.jpg"
	//
	//lib.DownLoad(url, "download", lib.GetPicName(url))
}
