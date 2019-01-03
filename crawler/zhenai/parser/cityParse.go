package parser

import (
	"goExProject/crawler/engine"
	"regexp"
)

func GetUserListMsg(bytes []byte) engine.ParserResult {
	re := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	matchs := re.FindAllSubmatch(bytes, -1)
	var results engine.ParserResult
	for _, m := range matchs {
		results.Items = append(results.Items, string(m[2]))
		results.Requests = append(
			results.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(bytes []byte) engine.ParserResult {
					return GetUserProfile(bytes, string(m[2]))
				},
			})
	}
	return results
}
