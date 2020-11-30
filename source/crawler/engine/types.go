package engine

type Request struct {
	Url string
	ParseFunc func([]byte) ParsedResult
}

type ParsedResult struct {
	Requests []Request
	Items    []interface{}
}

func DefaultParseFunc(content []byte) ParsedResult {
	return ParsedResult{}
}