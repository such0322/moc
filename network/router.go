package network

type IRouter interface {
	GetName() string
	SetName(name string)
	Handle(req IRequest)
}

type BaseRouter struct {
	Name string
}

func (b BaseRouter) Handle(req IRequest) {}

func (b BaseRouter) GetName() string {
	return b.Name
}

func (b *BaseRouter) SetName(name string) {
	b.Name = name
}
