package request

type httpRequest struct {
	URL         string
	Method      string
	Headers     map[string]string
	QueryParams map[string]string
	Body        string
	TimeoutMs   int
}
