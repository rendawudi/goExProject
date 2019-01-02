package parser

import (
	"fmt"
	"goExProject/crawler/fetcher"
	"testing"
)

var testCity = []string{
	"毕节",
	"滨海新",
	"滨州",
}

var testUrl = []string{
	"http://www.zhenai.com/zhenghun/bijie",
	"http://www.zhenai.com/zhenghun/binhaixin",
	"http://www.zhenai.com/zhenghun/binzhou",
}

func TestCityListParse(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	results := GetCityList(contents)

	for m := range results.Items {
		fmt.Printf("Url: %s ; City: %s \n", results.Requests[m].Url, results.Items[m])
	}
}
