package parser

import (
	"regexp"
	"spider-local/engine"
)

func ParseCity(contents []byte) engine.ParseResult {
	var result engine.ParseResult
	//<a href="http://album.zhenai.com/u/1866103911" target="_blank"><img src="https://photo.zastatic.com/images/photo/466526/1866103911/1308554241883026.jpg?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="未婚董事长">
	re := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank"><img src="[^"]*" alt="([^"]+)">`)
	machs := re.FindAllSubmatch(contents, -1)
	for _, v := range machs {
		result.Items = append(result.Items, "userName:"+string(v[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(v[1]),
				ParseFunc: ParseProfile,
			})
	}
	return result
}
