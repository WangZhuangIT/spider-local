package parser

import (
	"log"
	"regexp"
	"spider/engine"
	"spider/model"
	"strconv"
	"strings"
)

func ParseProfile(contents []byte, url string) engine.ParseResult {
	var result engine.ParseResult
	item := formatItem(contents, url)
	if item == nil {
		return result
	}
	result.Items = append(result.Items, *item)
	return result
}

func formatItem(contents []byte, url string) *engine.Item {
	var user model.User
	item := new(engine.Item)
	//class="nickName">duck</h1>
	//class="id">ID：1570330383</div>
	//class="des f-cl">佛山 | 33 岁 | 高中及以下 | 未婚 | 179cm | 5001-8000 元<
	reName := regexp.MustCompile(`class="nickName"[^>]*>([^<]+)</h1>`)
	reId := regexp.MustCompile(`class="id"[^>]*>ID：([0-9]+)</div>`)
	reDetail := regexp.MustCompile(`class="des f-cl"[^>]*>([^<]+)<`)
	machsName := reName.FindSubmatch(contents)
	machsID := reId.FindSubmatch(contents)
	machsDetail := reDetail.FindSubmatch(contents)

	if len(machsName) < 2 || len(machsID) < 2 || len(machsDetail) < 2 {
		log.Printf("profile %v %v %v %v", url, len(machsName), len(machsID), len(machsDetail))
		return nil
	}
	user.Name = string(machsName[1])
	machsDetailArr := strings.Split(strings.Replace(string(machsDetail[1]), " ", "", -1), "|")

	if len(machsDetailArr) != 6 {
		log.Printf("profile %v %v", url, len(machsDetailArr))
		return nil
	}

	age, _ := strconv.Atoi(strings.Replace(machsDetailArr[1], "岁", "", -1))
	height, _ := strconv.Atoi(strings.Replace(machsDetailArr[4], "cm", "", -1))
	user.Address = machsDetailArr[0]
	user.Age = age
	user.Education = machsDetailArr[2]
	user.Married = machsDetailArr[3]
	user.Height = height
	user.Salary = strings.Replace(machsDetailArr[5], "元", "", -1)

	item.Payload = user
	item.Id = string(machsID[1])
	item.Type = "zhenai"
	item.Url = url
	return item

}
