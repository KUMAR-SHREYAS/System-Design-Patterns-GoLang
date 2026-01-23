package main

import "fmt"

func main() {
	auth := NewAuthHandler()
	authorization := NewAuthorizationHandler()
	rateLimit := NewRateLimitingHandler()
	validation := NewValidationHandler()
	businessLogic := NewBusinessHandler()

	auth.SetNext(authorization)
	authorization.SetNext(rateLimit)
	rateLimit.SetNext(validation)
	validation.SetNext(businessLogic)

	request := NewRequest("john", "ADMIN", 10, "{ \"data\": \"valid\" }")
	auth.Handle(*request)

	fmt.Println("\n--- Trying an invalid request ---")
	badRequest := NewRequest("", "USER", 150, "")
	auth.Handle(*badRequest)
}
