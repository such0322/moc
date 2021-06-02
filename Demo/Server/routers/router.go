package routers

import (
	"moc/network"
)

const (
	Ping = uint32(iota)
	Gate
	Hello
)

func RouterInit(s *network.Server) {
	s.AddRouter(Ping, "PingRouter", &PingRouter{})
	s.AddRouter(Gate, "GateRouter", &GateRouter{})
	s.AddRouter(Hello, "HelloRouter", &HelloRouter{})
}
