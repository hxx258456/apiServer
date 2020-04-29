package main

import (
	"encoding/json"
	"github.com/Unknwon/goconfig"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
	"webServer/models"
	"webServer/mqttConn"
	"webServer/routers"
	"webServer/utils"
)

func main () {
	exist, err := utils.PathExists("./request.log")
	if err != nil {
		utils.MyLog.WithFields(utils.ServerFields).WithFields(logrus.Fields{
			"action": "初始化request log文件",
			"result": "错误",
		}).Panic(err)
	}
	var reqLogFile *os.File
	if exist {
		reqLogFile, _ = os.OpenFile("./request.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
		restartTime := time.Now().Format("2006-01-02 15:04:05")
		restartLog :=strings.Join([]string{"\n","============",restartTime," server restart ============","\n","\n"},"")
		reqLogFile.Write([]byte(restartLog))
	} else {
		reqLogFile, _ = os.Create("./request.log")
	}
	//gin.DefaultWriter = io.MultiWriter(reqLogFile) //重定向log输出至文件
	message := struct {
		Msg string `json:"msg"`
	}{
		Msg:"hello world",
	}
	topic := "test1"
	utils.MyLog.WithFields(utils.MqttFields).WithFields(logrus.Fields{
		"action": "publish",
		"topic" : topic,
	}).Info(message)
	result,_ := json.Marshal(message)
	token := mqttConn.MqttConn.Publish("test1", 1, false, result)
	if token.Wait() && token.Error() != nil {
		utils.MyLog.WithFields(utils.MqttFields).WithFields(logrus.Fields{
			"action": "publish",
			"topic": topic,
		}).Error(token.Error())
	}
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		utils.MyLog.WithFields(utils.ServerFields).WithFields(logrus.Fields{
			"action": "读取配置文件",
			"result": "失败",
		}).Panic(err)
	}
	section, err := cfg.GetSection("server")
	if err != nil {
		utils.MyLog.WithFields(utils.ServerFields).WithFields(logrus.Fields{
			"action": "读取server配置信息",
			"result": "失败",
		}).Panic(err)
	}
	server := routers.InitServer()
	server.Run(":"+section["port"])
	defer func() {
		models.DB.Close()
		mqttConn.MqttConn.Disconnect(250)
		reqLogFile.Close()
		utils.ServerLogFile.Close()
	}()
}
