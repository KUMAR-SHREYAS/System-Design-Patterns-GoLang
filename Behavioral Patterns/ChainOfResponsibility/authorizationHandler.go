package main

import (
	"fmt"
)

type AuthorizationHandler struct {
	BaseHandler
}

func NewAuthorizationHandler() *AuthorizationHandler {
	return &AuthorizationHandler{}
}

func (a *AuthorizationHandler) Handle(request Request) {
	if !(request.role == "ADMIN") {
		fmt.Println("AuthorizationHandler: ❌ User not authenticated.")
		return
	}
	fmt.Println("AuthorizationHandler: ✅ Authorized.")
	a.forward(request)
}
