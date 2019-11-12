package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Worker chan Request

func NilRequestFunc([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Id      string
	Url     string
	Type    string
	Payload interface{}
}
