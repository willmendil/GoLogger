package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Log struct {
	printLog   *log.Logger
	traceLog   *log.Logger
	debugLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	fatalLog   *log.Logger
}

func (l *Log) Print(message string) {
	l.printLog.Output(2, message)
}

func (l *Log) Printf(format string, v ...interface{}) {
	l.printLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Log) Trace(message string) {
	l.traceLog.Output(2, message)
}

func (l *Log) Tracef(format string, v ...interface{}) {
	l.traceLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Log) Debug(message string) {
	l.debugLog.Output(2, message)
}

func (l *Log) Debugf(format string, v ...interface{}) {
	l.debugLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Log) Info(message string) {
	l.infoLog.Output(2, message)
}

func (l *Log) Infof(format string, v ...interface{}) {
	l.infoLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Log) Warning(message string) {
	l.warningLog.Output(2, message)
}

func (l *Log) Warningf(format string, v ...interface{}) {
	l.warningLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Log) Error(message string) {
	l.errorLog.Output(2, message)
}

func (l *Log) Errorf(format string, v ...interface{}) {
	l.errorLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Log) Fatal(message string) {
	l.fatalLog.Output(2, message)
	os.Exit(1)
}

func (l *Log) Fatalf(format string, v ...interface{}) {
	l.fatalLog.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// New creates an instance of Log.
// ex: log := logger.New("log.txt")
func New(outfile ...string) Log {
	l := Log{}
	writer := io.Writer(os.Stdout)
	if len(outfile) == 1 {
		logFile, err := os.OpenFile(outfile[0], os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		writer = io.MultiWriter(os.Stdout, logFile)
	}
	l.printLog = log.New(writer, "[Print] - ", log.Ldate|log.Ltime|log.Lshortfile)
	l.traceLog = log.New(writer, "[Trace] - ", log.Ldate|log.Ltime|log.Lshortfile)
	l.debugLog = log.New(writer, "[Debug] - ", log.Ldate|log.Ltime|log.Lshortfile)
	l.infoLog = log.New(writer, "[Info] - ", log.Ldate|log.Ltime|log.Lshortfile)
	l.warningLog = log.New(writer, "[Warning] - ", log.Ldate|log.Ltime|log.Lshortfile)
	l.errorLog = log.New(writer, "[Error] - ", log.Ldate|log.Ltime|log.Lshortfile)
	l.fatalLog = log.New(writer, "[Fatal] - ", log.Ldate|log.Ltime|log.Lshortfile)

	return l
}
