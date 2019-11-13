package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"spider/engine"
	"spider/model"
	"testing"
)

func TestSave(t *testing.T) {

	user := engine.Item{
		Id:   "10086",
		Url:  "www",
		Type: "zhenai",
		Payload: model.User{
			Name:      "wangdazhuang",
			Age:       18,
			Address:   "宁静的许庄村",
			Education: "本科毕业哦",
			Married:   "未婚",
			Salary:    "10086",
			Height:    65,
		},
	}

	// TODO: try to start up es search
	// here using docker go client,启动es
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		t.Fatal(err)
	}

	const index = "dating_test"
	err = save(user, client, index)

	if err != nil {
		t.Fatal(err)
	}

	result, err := client.Get().
		Index("dating_profile").
		Type(user.Type).
		Id(user.Id).
		Do(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	var searchUser engine.Item

	err = json.Unmarshal(*result.Source, &searchUser)
	if err != nil {
		t.Fatal(err)
	}
	searchUser.Payload, err = model.FormatJson(searchUser.Payload)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user == searchUser)

}
