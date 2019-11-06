package fetcher

import "testing"

func TestFetch(t *testing.T) {
	contents, err := Fetch("http://album.zhenai.com/u/1445144756")
	t.Log(string(contents), err)
}
