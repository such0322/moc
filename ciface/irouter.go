package ciface

type IRouter interface {
	Handle(req IRequest)
}
