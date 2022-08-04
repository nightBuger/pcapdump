package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
)

var Logger *log.Logger

func init() {
	loggerInit()
}

func CheckFmt(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(201)
	}
}

func loggerInit() {
	Logger = log.New()
	logfile, err := os.OpenFile("./goserver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	CheckFmt(err)
	writer := io.MultiWriter(os.Stdout, logfile)
	Logger.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05,000000",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "[" + frame.Function + "]", "[" + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + "]"
		},
	})
	Logger.SetReportCaller(true)
	Logger.SetOutput(writer)
	Logger.SetLevel(log.DebugLevel)
}
