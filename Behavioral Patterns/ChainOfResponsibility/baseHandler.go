package main

type BaseHandler struct {
	next RequestHandler
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// By implementing SetNextv in an abtract class
// every concrete handler can focus solely on its logic and delegate to forward(request) when needed.
func (b *BaseHandler) SetNext(next RequestHandler) {
	b.next = next
}

func (b *BaseHandler) forward(request Request) {
	if b.next != nil {
		b.next.Handle(request)
	}
}
