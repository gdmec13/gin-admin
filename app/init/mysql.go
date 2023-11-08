package init

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"simple-cloud-storage/app/global"
	"simple-cloud-storage/app/model"
	"simple-cloud-storage/pkg/util"
)

func SetupMysql() {
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

func GenerateTable() {
	flagFile := "flag.log"
	if ok, _ := util.PathExists(flagFile); !ok {
		createDBTables()
		if ok := util.GenerateFile(flagFile, ""); ok != nil {
			global.APP_LOG.Error("写入flag.log失败：", ok)
		}
	}
}

func createDBTables() {
	db := global.APP_DB

	err := db.AutoMigrate(model.MyFile{},
		model.User{},
		model.FileFolder{},
		model.FileStore{},
		model.Share{})

	if err != nil {
		global.APP_LOG.Error("init create table error", err)
		os.Exit(0)
	} else {
		global.APP_LOG.Debug("init create table success")
	}
}
