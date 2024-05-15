package types

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

type AppLogT struct {
	logger *log.Logger
	level  string
	format string
}

type sourceJson struct {
	Level    string `json:"level"`
	File     string `json:"file"`
	Line     int    `json:"line"`
	Message  string `json:"message"`
	DataTime string `json:"data"`
}

const (
	AppLogTypeText string = "text" // text format
	AppLogTypeJson string = "json" // json format

	AppLogLevelTrace string = "Trace" // 追踪
	AppLogLevelDebug string = "Debug" // 调试
	AppLogLevelInfo  string = "Info"  // 信息
	AppLogLevelWarn  string = "Warn"  // 警告
	AppLogLevelError string = "Error" // 错误
	AppLogLevelFatal string = "Fatal" // 严重
)

const (
	appLogLevelIntTrace int = iota + 1
	appLogLevelIntDebug
	appLogLevelIntInfo
	appLogLevelIntWarn
	appLogLevelIntError
	appLogLevelIntFatal
)

type AppLogI interface {
	SetLevel(level string)
	Trace(msgArr ...any)
	Debug(msgArr ...any)
	Info(msgArr ...any)
	Warn(msgArr ...any)
	Error(msgArr ...any)
	Fatal(msgArr ...any)
	print(level string, msgArr ...any)
}

func NewAppLog(level string, outputFile string, outputConsole bool, format string) *AppLogT {

	writer, ok := setOutput(outputFile, outputConsole)
	if !ok {
		log.Fatal("error: log output set failute")
		return nil
	}

	if format == "" {
		format = AppLogTypeText
	}

	return &AppLogT{
		logger: log.New(writer, "", log.LstdFlags|log.Lshortfile),
		level:  level,
		format: format,
	}
}

func setOutput(filename string, outputConsole bool) (io.Writer, bool) {
	if filename != "" {
		logFilePath := path.Join("logs", filename)
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Fatal(err)
			return nil, false
		}
		multiWriter := io.MultiWriter(file, os.Stdout)

		return multiWriter, true
	}

	if outputConsole {
		return os.Stdout, true
	}

	return nil, false
}

func (a *AppLogT) SetLevel(level string) {
	a.level = level
}

func (a *AppLogT) SetFormatType(format string) {
	a.format = format
}

func (a *AppLogT) Trace(msgArr ...any) {
	a.logger.SetPrefix("[Trace] ")
	a.print(AppLogLevelTrace, msgArr...)
}

func (a *AppLogT) Info(msgArr ...any) {
	a.logger.SetPrefix("[Info] ")
	a.print(AppLogLevelInfo, msgArr...)
}

func (a *AppLogT) Debug(msgArr ...any) {
	a.logger.SetPrefix("[Debug] ")
	a.print(AppLogLevelDebug, msgArr...)
}

func (a *AppLogT) Warn(msgArr ...any) {
	a.logger.SetPrefix("[Warn] ")
	a.print(AppLogLevelWarn, msgArr...)
}

func (a *AppLogT) Error(msgArr ...any) {
	a.logger.SetPrefix("[Error] ")
	a.print(AppLogLevelError, msgArr...)
}

func (a *AppLogT) Fatal(msgArr ...any) {
	a.logger.SetPrefix("[Fatal] ")
	a.print(AppLogLevelError, msgArr...)
	os.Exit(1)
}

func levelStringToInt(level string) int {
	switch level {
	case AppLogLevelTrace:
		return appLogLevelIntTrace
	case AppLogLevelDebug:
		return appLogLevelIntDebug
	case AppLogLevelInfo:
		return appLogLevelIntInfo
	case AppLogLevelWarn:
		return appLogLevelIntWarn
	case AppLogLevelError:
		return appLogLevelIntError
	case AppLogLevelFatal:
		return appLogLevelIntFatal
	default:
		return appLogLevelIntTrace
	}
}

func (a *AppLogT) print(level string, msgArr ...any) {
	if levelStringToInt(a.level) > levelStringToInt(level) {
		return
	}

	if a.format == AppLogTypeJson {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}

		// log.Println(a.logger.Flags())
		// if a.logger.Flags()&(log.Lshortfile|log.Llongfile) != 0 {
		// } else {
		// }

		// Lshortfile
		// file = filepath.Base(file)

		// Llongfile
		// file
		var sourceJson = sourceJson{
			Message:  fmt.Sprint(msgArr...),
			DataTime: time.Now().Format("2006-01-02 15:04:05"),
			Level:    level,
			File:     file,
			Line:     line,
		}
		jsonData, _ := json.Marshal(sourceJson)

		a.logger.Writer().Write(append(jsonData, '\n'))
	} else {
		s := fmt.Sprint(msgArr...)
		a.logger.Output(3, s)
	}
}
