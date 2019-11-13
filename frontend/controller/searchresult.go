package controller

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"regexp"
	"spider/engine"
	"spider/frontend/model"
	"spider/frontend/view"
	"strconv"
	"strings"
)

// TODO
// add start page
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := s.GetSearchResult(q, from)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = s.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s SearchResultHandler) GetSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := s.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	result.Query = q
	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
