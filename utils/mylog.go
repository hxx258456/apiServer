package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var MyLog *logrus.Logger
var ServerLogFile *os.File

var DbFields = logrus.Fields{
	"type": "Mysql",
}
var MqttFields = logrus.Fields{
	"type": "MQTT",
}
var ServerFields = logrus.Fields{
	"type": "Server",
}

func init () {
	MyLog = logrus.New()
	exist, err := PathExists("./server.log")
	if err != nil {
		MyLog.WithFields(ServerFields).WithFields(logrus.Fields{
			"action": "初始化server log文件",
			"result": "错误",
		}).Panic(err)
	}
	if exist {
		ServerLogFile, _ = os.OpenFile("./server.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
		restartTime := time.Now().Format("2006-01-02 15:04:05")
		restartLog :=strings.Join([]string{"\n","============",restartTime," server restart ============","\n","\n"},"")
		ServerLogFile.Write([]byte(restartLog))
	} else {
		ServerLogFile, _ = os.Create("./server.log")
	}
	MyLog.SetOutput(ServerLogFile)
}