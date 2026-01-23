package main

type RequestHandler interface {
	SetNext(next RequestHandler)
	Handle(request Request)
}
