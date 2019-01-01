package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Status Code :", resp.StatusCode)
		return
	}
	
	encode := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, encode.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		panic(err)
	}

	printCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(bytes []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matchs := re.FindAllSubmatch(bytes, -1)
	for _, m := range matchs {
		fmt.Printf("City: %s  URL: %s", m[2], m[1])
		fmt.Println()
	}
}