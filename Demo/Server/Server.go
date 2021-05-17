package main

import (
	"moc/Demo/Server/routers"
	"moc/network"
)

func main() {

	s := network.NewServer()
	//routersInit := routers.RouterInit(s)
	routers.RouterInit(s)

	s.Serve()
}
