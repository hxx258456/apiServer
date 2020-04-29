package models

import (
	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"webServer/utils"
)

var DB *gorm.DB

func init () {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil{
		utils.MyLog.WithFields(utils.ServerFields).WithFields(logrus.Fields{
			"action": "读取配置文件",
			"result": "失败",
		}).Panic(err)
	}
	section, err := cfg.GetSection("mysql")
	if err != nil {
		utils.MyLog.WithFields(utils.DbFields).WithFields(logrus.Fields{
			"action": "获取mysql配置信息错误",
			"result": "失败",
		}).Panic(err)
	}
	username := section["username"] //"root"
	password := section["password"] //"root"
	dataname := section["dataname"] //"device"
	port :=  section["port"] //"3306"
	host :=  section["host"]//"127.0.0.1"
	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?parseTime=True&loc=Local"
	utils.MyLog.WithFields(utils.DbFields).WithFields(logrus.Fields{
		"action": "conn db",
		"host": host,
		"port": port,
		"username":username,
		"password":password,
		"db": dataname,
	}).Info("success")
	db, err := gorm.Open("mysql", dns)
	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "device_" + defaultTableName
	}
	if err != nil {
		utils.MyLog.WithFields(utils.DbFields).WithFields(logrus.Fields{
			"action": "conn db",
			"host": host,
			"port": port,
			"username":username,
			"password":password,
			"db": dataname,
		}).Error(err.Error())
	}
	DB = db
	db.LogMode(false)
	db.AutoMigrate(&Conf{}) // 自动迁移，自动创建表
}
