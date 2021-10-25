package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
}

var (
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init() {
	initial(os.Stdout, os.Stdout, os.Stderr)
}

func New() (*Logger, error) {

	l := &Logger{
	}

	return l, nil
}

func (l *Logger) Init(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	l.Info = log.New(infoHandle, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Warning = log.New(warningHandle, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Error = log.New(errorHandle, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

}

func initial(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Info = log.New(infoHandle, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

}
