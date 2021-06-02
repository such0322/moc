package network

import (
	"fmt"
	"moc/logs"
)

type IMsgHandle interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgID uint32, router IRouter)
}

type MsgHandle struct {
	Apis map[uint32]IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]IRouter),
	}
}

func (mh *MsgHandle) DoMsgHandler(request IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID =", request.GetMsgID(), " is not found!")
		return
	}
	logs.Debug("send to:", handler.GetName())
	handler.Handle(request)
}

func (mh *MsgHandle) AddRouter(msgID uint32, router IRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		fmt.Println("api repeat, msgID =", msgID)
	}
	mh.Apis[msgID] = router
	fmt.Println("add api msgID =", msgID)
}
