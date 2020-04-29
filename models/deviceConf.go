package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"webServer/utils"
)

type Conf struct {
	gorm.Model
	DeviceName string `gorm:"not null;size:255;column:devicename;unique_index"`
	ValveAngle int `gorm:"not null;column:valve;"`
	HuntingRange int `gorm:"not null;column:hunting""`
	Note string `gorm:"not null;column:note"`
	Test string `gorm:"not null;column:test"`
	Share string `gorm:"not null;column:share"`
} // 设备配置表

func GetDeviceConf (deviceId int) (*Conf, error){
	result := new(Conf)
	if err := DB.Where("devicename = ?",  deviceId).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllDeviceConf () ([]Conf,error) {
	list := make([]Conf, 0)
	if err := DB.Find(&list).Error; err != nil {
		return nil,err
	}
	return list, nil
}

func InsertTest (conf *Conf) (error){
	//conf := &Conf{
	//	Model:        gorm.Model{},
	//	DeviceName:   "",
	//	ValveAngle:   0,
	//	HuntingRange: 0,
	//}
	DB.NewRecord(conf)
	if err := DB.Omit("devicename", "valve", "hunting").Create(conf).Error; err != nil {
		utils.MyLog.WithFields(utils.DbFields).WithFields(logrus.Fields{
			"action": "create data",
			"data": fmt.Sprintf("%+v",conf),
		}).Error(err.Error())
		return err
	}
	DB.NewRecord(conf)
	return nil
}

func UpdateConf (devicename string, updateVal map[string]interface{}) error {
	if err := DB.Model(&Conf{}).Where("device = ?", devicename).Update(updateVal).Error; err != nil {
		return err
	}
	return nil
}
