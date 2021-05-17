package network


type IRequest interface {
	GetMsgID() uint32
	GetData() []byte
}

type Request struct {
	Msg IMessage
}

func (r *Request) GetMsgID() uint32 {
	return r.Msg.GetID()
}

func (r *Request) GetData() []byte {
	return r.Msg.GetData()
}
