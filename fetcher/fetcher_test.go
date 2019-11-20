package fetcher

import "testing"

func TestFetch(t *testing.T) {
	contents, err := Fetch("https://album.zhenai.com/u/1059758035", "local")
	t.Log(string(contents), err)
}
