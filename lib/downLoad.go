package lib

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func GetPicName(url string) string {
	restr := "(\\w|\\d|_)*.(jpg|png|jpeg)"
	reg, _ := regexp.Compile(restr)
	name := reg.FindStringSubmatch(url)[0]
	return name
}

// to get the Image name

func DownLoad(url string, outPut string) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Referer", "https://weibo.com/")
	// REFERER MUST BE WEIBO!
	request.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"103\", \".Not/A)Brand\";v=\"99\"")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// you can change your User-Agent here

	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error! " + resp.Status)
	}
	defer resp.Body.Close()

	res, _ := io.ReadAll(resp.Body)

	out, err := os.Create(outPut)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(out, bytes.NewReader(res))
	fmt.Println("Download " + outPut + " successfully!")
	return err
}
