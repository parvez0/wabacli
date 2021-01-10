package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init() {
	level, err := logrus.ParseLevel(os.Getenv("WABA_LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)
	log.Formatter =  &logrus.TextFormatter{
		DisableColors: false,
		DisableTimestamp: true,
		DisableLevelTruncation: true,
		FullTimestamp: false,
		DisableQuote: true,
	}
}

func Info(msgs ...interface{}) {
	log.Info(msgs ...)
}

func Warn(msgs ...interface{}) {
	log.Warn(msgs ...)
}

func Debug(msgs ...interface{}) {
	log.Debug(msgs ...)
}

func Error(msgs ...interface{}) {
	log.Error(msgs ...)
	os.Exit(1)
}

func Fatal(msgs ...interface{}) {
	log.Fatal(msgs ...)
}

func Panic(msgs ...interface{}) {
	log.Panic(msgs ...)
}

func WithFieldsInfo(fields map[string]interface{}, msg string)  {
	log.WithFields(fields).Info(msg)
}

func WithFieldsError(fields map[string]interface{}, msg string)  {
	log.WithFields(fields).Error(msg)
}

func WithFieldsWarn(fields map[string]interface{}, msg string)  {
	log.WithFields(fields).Debug(msg)
}