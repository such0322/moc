package network

import (
	"fmt"
	"moc/utils"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int

	MsgHandle IMsgHandle
}

func (s *Server) AddRouter(msgId uint32, router IRouter) {
	s.MsgHandle.AddRouter(msgId, router)
	fmt.Println("Add Router Succ!!")
}

func NewServer() *Server {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.IP,
		Port:      utils.GlobalObject.Port,
		MsgHandle: NewMsgHandle(),
	}
	return s
}

func (s *Server) Start() {
	fmt.Println("[START]moc server start! ServerName: ", s.Name, "IP: ", s.IP, "Port: ", s.Port)

	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve Tcp Add err:", err)
		return
	}
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen Tcp err:", err)
		return
	}

	var cid uint32
	cid = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept Tcp err: ", err)
			continue
		}

		dealConn := NewConnection(conn, cid, s.MsgHandle)
		cid++

		go dealConn.Start()
	}
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}
