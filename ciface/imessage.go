package ciface

type IMessage interface {
	GetID() uint32
	GetDataLen() uint32
	GetData() []byte
}
