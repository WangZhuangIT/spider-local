package parser

import (
	"regexp"
	"spider/engine"
)

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	var result engine.ParseResult
	////<a href="http://www.zhenai.com/zhenghun/bangbu" data-v-5e16505f>蚌埠</a>
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	machs := re.FindAllSubmatch(contents, -1)
	for _, v := range machs {
		//result.Items = append(result.Items, "City:"+string(v[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(v[1]),
				ParseFunc: ParseCity,
			})
	}
	return result
}
