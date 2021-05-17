package ciface

type IRequest interface {
	GetMsgID() uint32
	GetData() []byte
}
