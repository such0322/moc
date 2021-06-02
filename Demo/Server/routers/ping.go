package routers

import (
	"fmt"
	"moc/network"
)

type PingRouter struct {
	network.BaseRouter
}

func (r *PingRouter) Handle(request network.IRequest) {
	fmt.Println("pingpingpingpingping")
}
