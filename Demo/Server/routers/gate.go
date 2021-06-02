package routers

import (
	"fmt"
	"moc/network"
)

type GateRouter struct {
	network.BaseRouter
}

func (r GateRouter) Handle(req network.IRequest) {
	fmt.Println("this is gate router")
}
