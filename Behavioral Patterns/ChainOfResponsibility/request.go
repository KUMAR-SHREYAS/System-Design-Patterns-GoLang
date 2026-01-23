package main

type Request struct {
	user         string
	role         string
	requestCount int
	payload      string
}

func NewRequest(user, role string, requestCount int, payload string) *Request {
	return &Request{
		user:         user,
		role:         role,
		requestCount: requestCount,
		payload:      payload,
	}
}
