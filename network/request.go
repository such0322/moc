package network

import "moc/ciface"

type Request struct {
	Msg ciface.IMessage
}

func (r *Request) GetMsgID() uint32 {
	return r.Msg.GetID()
}

func (r *Request) GetData() []byte {
	return r.Msg.GetData()
}
