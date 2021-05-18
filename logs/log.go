package logs

import (
	"fmt"
	"os"
	"strings"
)

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

const (
	AdapterConsole = "console"
	AdapterFile    = "file"
)

type Logger interface {
	Init() error
	WriteMsg(msg string) error
	Destroy()
	Flush()
}

type MocLogger struct {
	logger Logger
	level  int
}

var mocLogger = NewLogger()

func NewLogger() *MocLogger {
	ml := &MocLogger{
		level: LevelDebug,
	}
	ml.setLogger(AdapterConsole)
	return ml
}

func (ml *MocLogger) setLogger(adapterName string) error {
	var lg Logger
	switch adapterName {
	case AdapterConsole:
		lg = NewConsoleWrite()
	case AdapterFile:
		lg = NewFileWriter()
	}
	if lg == nil {
		return fmt.Errorf("logs:unknown logger")
	}
	ml.logger = lg
	lg.Init()
	return nil
}

func SetLogger(adapterName string) error {
	return mocLogger.setLogger(adapterName)
}

func Reset() {
	mocLogger.Reset()
}

func (ml *MocLogger) Reset() {
	ml.flush()
	ml.logger.Destroy()
	ml.logger = nil
}

func (ml *MocLogger) flush() {
	ml.logger.Flush()
}

func (ml *MocLogger) writeMsg(msg string) {
	ml.writeToLoggers(msg)
}

func (ml *MocLogger) writeToLoggers(msg string) {
	err := ml.logger.WriteMsg(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to WriteMsg to logger:", err)
	}
}

func (ml *MocLogger) Debug(format string, v ...interface{}) {
	if LevelDebug > ml.level {
		return
	}
	ml.writeMsg(format)
}

func Debug(f interface{}, v ...interface{}) {
	mocLogger.Debug(formatLog(f, v...))
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if !strings.Contains(msg, "%") {
			// do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
