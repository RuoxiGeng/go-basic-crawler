package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.7799520.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url : string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
