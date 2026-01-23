package main

import "fmt"

type AuthHandler struct {
	BaseHandler
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) Handle(request Request) {
	if request.user == "" {
		fmt.Println("AuthHandler: ❌ User not authenticated.")
		return
	}
	fmt.Println("AuthHandler: ✅ Authenticated.")
	a.forward(request)
}
