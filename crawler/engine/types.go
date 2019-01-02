package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Items    []interface{}
	Requests []Request
}

func NilParseFunc([]byte) ParserResult {
	return ParserResult{}
}
