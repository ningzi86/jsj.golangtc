package engine

type Request struct {
	Url        string
	ParserFunc func(contents []byte) ParseResult
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type ParseResult struct {
	Items   []Item
	Request []Request
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
