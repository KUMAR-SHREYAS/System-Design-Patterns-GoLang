package request

type httpRequestBuilder struct {
	url         string
	method      string
	headers     map[string]string
	queryParams map[string]string
	body        string
	timeoutMs   int
}

func NewHttpRequestBuilder(url string) *httpRequestBuilder {
	return &httpRequestBuilder{
		url:         url,
		method:      "GET",
		headers:     make(map[string]string),
		queryParams: make(map[string]string),
		timeoutMs:   30000,
	}
}

func (b *httpRequestBuilder) Method(method string) *httpRequestBuilder {
	b.method = method
	return b
}

func (b *httpRequestBuilder) AddHeader(key, value string) *httpRequestBuilder {
	b.headers[key] = value
	return b
}

func (b *httpRequestBuilder) AddQueryParams(key, value string) *httpRequestBuilder {
	b.queryParams[key] = value
	return b
}

func (b *httpRequestBuilder) Body(body string) *httpRequestBuilder {
	b.body = body
	return b
}

func (b *httpRequestBuilder) Timeout(ms int) *httpRequestBuilder {
	b.timeoutMs = ms
	return b
}

func (b *httpRequestBuilder) Build() httpRequest {
	return httpRequest{
		URL:         b.url,
		Method:      b.method,
		Headers:     cloneMap(b.headers),
		QueryParams: cloneMap(b.queryParams),
		Body:        b.body,
		TimeoutMs:   b.timeoutMs,
	}
}
