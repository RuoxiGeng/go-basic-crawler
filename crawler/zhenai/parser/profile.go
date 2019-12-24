package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<span class="age s1">(\d+)岁</span>`)
var marriageRe  = regexp.MustCompile(`<span class="marrystatus">([^<]+)</span>`)
var heightRe = regexp.MustCompile(`<span class="height">(\d+)cm</span>`)
var genderRe  = regexp.MustCompile(`<span>([^<]+) </span>`)
var educationRe  = regexp.MustCompile(`<span class="education">([^<]+)</span>`)
var incomeRe  = regexp.MustCompile(`<li>收入：<span>([^<]+)</span></li>`)
var occupationRe  = regexp.MustCompile(`<li>职业：<span>([^<]+)</span></li>`)
var hukouRe  = regexp.MustCompile(`<li>籍贯：<span>([^<]+)</span></li>`)
var xingzuoRe  = regexp.MustCompile(`<li>星座：<span>([^<]+)</span></li>`)
var idUrlRe = regexp.MustCompile(`<meta name="mobile-agent" content="format=xhtml; url=http://m.7799520.com/user/([\d]+).html">`)

func ParseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	profile.Marriage = extractString(contents, marriageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Education = extractString(contents, educationRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type: "zhenai",
				Id: extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}