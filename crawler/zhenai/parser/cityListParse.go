package parser

import (
	"goExProject/crawler/engine"
	"regexp"
)

func GetCityList(bytes []byte) engine.ParserResult {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matchs := re.FindAllSubmatch(bytes, -1)
	var results engine.ParserResult
	for _, m := range matchs {
		results.Items = append(results.Items, string(m[2]))
		results.Requests = append(
			results.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: GetUserMsg,
			})
	}
	return results
}
