package parser

import (
	"regexp"
	"spider/engine"
)

var (
	//<a href="http://album.zhenai.com/u/1866103911" target="_blank"><img src="https://photo.zastatic.com/images/photo/466526/1866103911/1308554241883026.jpg?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="未婚董事长">
	regProfile = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank"><img src="[^"]*" alt="([^"]+)">`)
	//<a href="http://www.zhenai.com/zhenghun/shanghai/2">下一页</a>
	regCityUrl = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	var result engine.ParseResult
	machs := regProfile.FindAllSubmatch(contents, -1)
	for _, v := range machs {
		url := string(v[1])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       url,
				ParseFunc: ProfileParse(),
			})
	}

	machs = regCityUrl.FindAllSubmatch(contents, -1)

	for _, v := range machs {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(v[1]),
				ParseFunc: ParseCity,
			})
	}
	return result
}

func ProfileParse() engine.ParseFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url)
	}
}
