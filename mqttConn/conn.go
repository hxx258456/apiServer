package mqttConn

import (
	"github.com/Unknwon/goconfig"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"webServer/utils"
)

var MqttConn mqtt.Client

func init () {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil{
		utils.MyLog.WithFields(utils.ServerFields).WithFields(logrus.Fields{
			"action": "读取配置文件",
			"result": "失败",
		}).Panic(err)
	}
	section, err := cfg.GetSection("mqttBroker")
	if err != nil {
		utils.MyLog.WithFields(utils.DbFields).WithFields(logrus.Fields{
			"action": "获取mysql配置信息错误",
			"result": "失败",
		}).Panic(err)
	}
	user := section["user"]
	password := section["password"]
	port :=  section["port"]
	ip :=  section["ip"]
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + ip + ":" + port).SetUsername(user).SetPassword(password)
	utils.MyLog.WithFields(utils.MqttFields).WithFields(logrus.Fields{
		"ip": ip,
		"port": port,
		"user": user,
		"password": password,
	}).Info("start conn")
	MqttConn = mqtt.NewClient(opts)
	if token := MqttConn.Connect(); token.Wait() && token.Error() != nil {
		utils.MyLog.WithFields(utils.MqttFields).WithFields(logrus.Fields{
			"host": ip,
			"port": port,
			"user": user,
			"password": password,
		}).Panic(token.Error())
	}
	utils.MyLog.WithFields(utils.MqttFields).WithFields(logrus.Fields{
		"host": ip,
		"port": port,
		"user": user,
		"password": password,
	}).Info("conn success")
	if token := MqttConn.Subscribe("test1", 0, func(client mqtt.Client, message mqtt.Message) {
		utils.MyLog.WithFields(utils.MqttFields).WithFields(logrus.Fields{
			"action" : "Received",
			"topic" : message.Topic(),
		}).Info(string(message.Payload()))
	}); token.Wait() && token.Error() != nil {
		utils.MyLog.WithFields(utils.MqttFields).WithFields(logrus.Fields{
			"action" : "Received Error",
		}).Error(token.Error())
	}
}
