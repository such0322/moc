package routers

import (
	"fmt"
	"moc/network"
)

const (
	Ping = uint32(iota)
	Hello
)

type PingRouter struct {
	network.BaseRouter
}

func (r *PingRouter) Handle(request network.IRequest) {
	fmt.Println("pingpingpingpingping")
}

type HelloRouter struct {
	network.BaseRouter
}

func RouterInit(s network.IServer) {
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

}
