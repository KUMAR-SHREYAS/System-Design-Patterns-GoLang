package main

import "fmt"

type RateLimiting struct {
	BaseHandler
}

func NewRateLimitingHandler() *RateLimiting {
	return &RateLimiting{}
}

func (r *RateLimiting) Handle(request Request) {
	if request.requestCount >= 100 {
		fmt.Println("RateLimitHandler: ❌ Rate limit exceeded.")
		return
	}
	fmt.Println("RateLimitHandler: ✅ Within rate limit.")
	r.forward(request)
}
