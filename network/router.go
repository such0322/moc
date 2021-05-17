package network

type IRouter interface {
	Handle(req IRequest)
}

type BaseRouter struct {
}

func (b BaseRouter) Handle(req IRequest) {}
