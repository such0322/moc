package network

import (
	"fmt"
	"net"
)

type IConnection interface {
	Start()
	Stop()
}

type Connection struct {
	Conn   *net.TCPConn
	ConnID uint32

	MsgHandle IMsgHandle
}

func NewConnection(conn *net.TCPConn, cid uint32, msgHandle IMsgHandle) IConnection {
	c := &Connection{
		Conn:      conn,
		ConnID:    cid,
		MsgHandle: msgHandle,
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("conn read goroutine is running... ConnId =", c.ConnID)
	defer c.Stop()

	fmt.Println(c.MsgHandle)
	for {
		//读取数据流
		//TODO 这个应该是
		buf := make([]byte, 1024)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
			return
		}

		msg := NewMsgPackage(0, buf)

		req := &Request{
			Msg: msg,
		}

		//处理protobuf

		//交给逻辑层处理数据
		go c.MsgHandle.DoMsgHandler(req)

	}
}

func (c *Connection) StartWirter() {

}

func (c *Connection) Start() {
	go c.StartReader()
	go c.StartWirter()
}

func (c *Connection) Stop() {

}
