package global

import (
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"simple-cloud-storage/app/config"
)

var (
	APP_CONFIG config.Server
	APP_VP     *viper.Viper
	APP_LOG    *oplogging.Logger
	APP_DB     *gorm.DB
)

func CloseDb() {
	sqlDb, err := APP_DB.DB()
	if err == nil {
		sqlDb.Close()
	}
}
