package main

import (
	"fmt"
	"strings"
)

type ValidationHandler struct {
	BaseHandler
}

func NewValidationHandler() *ValidationHandler {
	return &ValidationHandler{}
}
func (v *ValidationHandler) Handle(request Request) {
	if request.payload == "" || strings.TrimSpace(request.payload) == "" {
		fmt.Println("ValidationHandler: ❌ Invalid payload.")
	}
	v.forward(request)
}
