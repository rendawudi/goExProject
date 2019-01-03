package fetcher

import (
	"bufio"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(Url string) ([]byte, error) {
	request, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(
		"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	request.Header.Add(
		"Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add(
		"Accept-Encoding", "gzip, deflate")
	request.Header.Add(
		"Accept-Language", "zh-CN,zh;q=0.9")

	var client http.Client
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error code: %v", resp.StatusCode)
	}

	encode := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, encode.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Error("fetcher encode error : %s", bytes)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
