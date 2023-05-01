package log

import (
	"github.com/Ryeom/soft-formula/internal"
	"github.com/labstack/echo/middleware"
	"github.com/op/go-logging"
	"os"
	"strings"
	"time"
)

const (
	ProjectName       = "soft-formula"
	ServerLogFileName = "server.log"
	AccessLogFileName = "access.log"
)

var Logger *logging.Logger

func InitializeApplicationLog() {
	var err error
	logPath := "/var/log" + ProjectName + "/"
	internal.CheckDirectoryPath(logPath)
	serverLogPath := logPath + ServerLogFileName
	internal.CheckFilePath(serverLogPath)
	accessLogPath := logPath + AccessLogFileName
	internal.CheckFilePath(accessLogPath)
	ServerLogDesc, err = os.OpenFile(serverLogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	AccessLogDesc, err = os.OpenFile(accessLogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	Logger = logging.MustGetLogger(ProjectName)
	back1 := logging.NewLogBackend(ServerLogDesc, "", 0)
	format := logging.MustStringFormatter(`%{color}%{time:0102 15:04:05.000} %{shortfunc:15s} ▶ %{level:.5s}%{color:reset} %{shortfile:15s} %{message}`)
	back1Formatter := logging.NewBackendFormatter(back1, format)
	back1Leveled := logging.AddModuleLevel(back1) //기본로그 외에 추가로그를 남기는 로직
	back1Leveled.SetLevel(logging.ERROR, "")      //추가로그의 로그 기본 레벨

	logging.SetBackend(back1Formatter)
	logging.SetLevel(logging.DEBUG, ProjectName)

	Logger.Info(banner)
}

var banner = strings.Repeat("░", 150) + "\n" + ProjectName + "\n" + strings.Repeat("░", 150)

var ServerLogDesc *os.File
var AccessLogDesc *os.File

func GetCustomLogConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `[${status} ${time_custom}] ${method}` +
			` ${uri} ${host} ${latency_human} ${error} ${bytes_in} ${bytes_out} ${remote_ip} ${header:Client-Ip}` + "\n",
		CustomTimeFormat: time.RFC3339,
		Output:           AccessLogDesc,
	}
}
