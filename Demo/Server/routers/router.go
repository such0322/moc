package routers

import (
	"fmt"
	"moc/ciface"
	"moc/network"
)

const (
	Ping = uint32(iota)
	Hello
)

type PingRouter struct {
	network.BaseRouter
}

func (r *PingRouter) Handle(request ciface.IRequest) {
	fmt.Println("pingpingpingpingping")
}

type HelloRouter struct {
	network.BaseRouter
}

func RouterInit(s ciface.IServer) {
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

}
