package parser

type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Request []Request
	Items []string
}

func NilFunc(body []byte) ParserResult {
	return ParserResult{}
}