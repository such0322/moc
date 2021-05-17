package network

import (
	"fmt"
	"moc/ciface"
)

type MsgHandle struct {
	Apis map[uint32]ciface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ciface.IRouter),
	}
}

func (mh *MsgHandle) DoMsgHandler(request ciface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID =", request.GetMsgID(), " is not found!")
		return
	}
	handler.Handle(request)
}

func (mh *MsgHandle) AddRouter(msgID uint32, router ciface.IRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		fmt.Println("api repeat, msgID =", msgID)
	}
	mh.Apis[msgID] = router
	fmt.Println("add api msgID =", msgID)
}
