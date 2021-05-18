package main

import (
	"moc/Demo/Server/routers"
	"moc/logs"
	"moc/network"
)

func main() {

	s := network.NewServer()
	logs.SetLogger(logs.AdapterFile)
	logs.Debug("test sample logger")
	//routersInit := routers.RouterInit(s)
	routers.RouterInit(s)

	s.Serve()
}
