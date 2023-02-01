package main

import (
	"SinaDownLoad/lib"
	"fmt"
)

func main() {
	path := "/Users/xiabee/Desktop/GitHub/gitpage/source/_posts/"

	files, err := lib.GetFileName(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(files)
	}

	//urlList, err := lib.UrlFinding(path)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, url := range urlList {
	//	fmt.Println(url)
	//}
	url := "https://tva1.sinaimg.cn/large/0084b03xgy1h58pm1gs0uj30y20iu0za.jpg"
	fmt.Println(lib.GetPicName(url))
	//lib.DownLoad(url, "./ss.jpg")
}
