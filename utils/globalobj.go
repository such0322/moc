package utils

import (
	"encoding/json"
	"io/ioutil"
)

type GlobalObj struct {
	Name string
	IP   string
	Port int

	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Name: "Moc Server App",
		IP:   "127.0.0.1",
		Port: 9100,

		Version: "0.0.1",
		MaxConn: 1000,

		MaxPackageSize: 4096,
	}
	//GlobalObject.Reload()
}

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
