package main

import (
	"moc/Demo/Server/routers"
	"moc/logs"
	"moc/network"
)

func main() {

	s := network.NewServer()
	logs.Init()
	logs.Debug("aasdfasdfasdfaa")
	//routersInit := routers.RouterInit(s)
	routers.RouterInit(s)

	s.Serve()
}
