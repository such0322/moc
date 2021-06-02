package network

import (
	"fmt"
	"io"
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
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			fmt.Println("read msg head error ", err)
			return
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			return
		}
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Println("read msg data error ", err)
				return
			}
		}
		msg.SetData(data)

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
