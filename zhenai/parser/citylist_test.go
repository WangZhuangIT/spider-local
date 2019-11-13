package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city_list_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d , but have %d ", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d , but have %d ", resultSize, len(result.Items))
	}

	for k, v := range result.Items {
		fmt.Println(v, k)
	}
}
