package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://www.7799520.com/user/[0-9]+.html)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.7799520.com/zhenghun/[^"]+)"`)
)


func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url : string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url : string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}
