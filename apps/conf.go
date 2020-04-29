package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
	"webServer/models"
)

func Test (context *gin.Context) {
	err := models.InsertTest(&models.Conf{
		Model:        gorm.Model{
			ID:        0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		},
		DeviceName:   "",
		ValveAngle:   0,
		HuntingRange: 0,
		Note:         "",
		Test:         "",
		Share:        "",
	})
	if err != nil {
		context.JSON(200, gin.H{
			"code": 0,
			"error": err.Error(),
		})
	} else {
		context.JSON(200, gin.H{
			"code": 1,
			"result": true,
		})
	}
}
