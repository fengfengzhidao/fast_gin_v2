package flags

import (
	"fast_gin/global"
	"fast_gin/models"
	"github.com/sirupsen/logrus"
)

func MigrateDB() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		logrus.Errorf("表结构迁移失败 %s", err)
		return
	}
	logrus.Infof("表结构迁移成功")
}
