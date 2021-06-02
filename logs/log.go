package logs

import "github.com/sirupsen/logrus"

var mocLogger = NewLogger()

type MocLogger struct {
	*logrus.Logger
}

func Init() {
	mocLogger.SetLevel(logrus.DebugLevel)

}

func NewLogger() *MocLogger {
	ml := &MocLogger{
		logrus.New(),
	}
	return ml
}

func Debug(arg ...interface{}) {
	mocLogger.Debug(arg...)
}
