package init

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"simple-cloud-storage/app/global"
)

func InitMysql() {
	dbConfig := global.APP_CONFIG.Mysql
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Uri + ")/" + dbConfig.Dbname + "?" + dbConfig.Config
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		global.APP_LOG.Error("Mysql 启动异常", err)
		os.Exit(0)
	} else {
		global.APP_LOG.Debug("Mysql 已启动")
		global.APP_DB = db
	}
}
