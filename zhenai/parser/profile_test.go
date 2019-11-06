package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents)

	for _, v := range result.Items {
		fmt.Println(v)
	}
}
