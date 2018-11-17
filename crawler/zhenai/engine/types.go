package engine

type Request struct {
	Url        string
	ParserFunc func(contents []byte) ParseResult
}

type ParseResult struct {
	Items   []interface{}
	Request []Request
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
