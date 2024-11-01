package logs

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
		PadLevelText:     false,
	})
	Logger.SetLevel(logrus.InfoLevel)
}

func LogError(err error) {
	pc, _, line, _ := runtime.Caller(2)
	function := runtime.FuncForPC(pc)

	Logger.WithFields(logrus.Fields{
		"location": function.Name(),
		"line":     line,
	}).Error(err.Error())
}
