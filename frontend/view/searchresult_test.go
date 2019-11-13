package view

import (
	"os"
	"spider/engine"
	"spider/frontend/model"
	common "spider/model"

	"testing"
)

func TestCreateSearchResultView(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Id:   "10086",
		Url:  "www",
		Type: "zhenai",
		Payload: common.User{
			Name:      "wangdazhuang",
			Age:       18,
			Address:   "宁静的许庄村",
			Education: "本科毕业哦",
			Married:   "未婚",
			Salary:    "10086",
			Height:    173,
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Fatal(err)
	}
}
