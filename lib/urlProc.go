package lib

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

func ReadLines(filename string) ([]string, error) {
	fd, err := os.Open(filename)
	var data []string
	if err != nil {
		log.Fatal(err)

	}
	defer fd.Close()
	buff := bufio.NewReader(fd)
	for {
		line, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		data = append(data, string(line))
	}
	return data, err
}

// Parse the text, return a slice

func GetPicName(url string) string {
	restr := "(\\w|\\d|_)*.(jpg|png|jpeg)"
	reg, _ := regexp.Compile(restr)
	name := reg.FindString(url)
	return name
}

// to get the Image name

func UrlFinding(filename string) ([]string, error) {
	regImg := "(img|image|imag)"
	regUrl := "(ht|f)tp(s?)\\:\\/\\/[0-9a-zA-Z]([-.\\w]*[0-9a-zA-Z])*(:(0-9)*)*(\\/?)([a-zA-Z0-9\\-\\.\\?\\,\\'\\/\\\\+&amp;%$#_]*)?"
	// regexp to find image URI and URL

	var urlList []string
	// image url list
	lines, err := ReadLines(filename)
	for _, line := range lines {
		regi := regexp.MustCompile(regImg)
		result := regi.FindAllStringSubmatch(line, -1)
		// locate the image URI, find all
		if result == nil {
			continue
		}
		regu := regexp.MustCompile(regUrl)
		if regu.FindString(line) != "" {
			url := regu.FindString(line)
			urlList = append(urlList, url, GetPicName(url))
		}
		// find the first url

	}
	return urlList, err
}

// to detect the pic's url
