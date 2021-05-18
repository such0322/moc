package logs

import "fmt"

type consoleWriter struct {
}

func NewConsoleWrite() Logger {
	cw := &consoleWriter{}
	return cw
}

func (c consoleWriter) Init() error {
	return nil
}

func (c consoleWriter) WriteMsg(msg string) error {
	fmt.Printf("%#v\n", msg)
	return nil
}

func (c consoleWriter) Destroy() {
}

func (c consoleWriter) Flush() {
}
