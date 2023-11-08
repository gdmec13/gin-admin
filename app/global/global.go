package global

import (
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"simple-cloud-storage/app/config"
)

var (
	APP_CONFIG config.Server
	APP_VP     *viper.Viper
	APP_LOG    *oplogging.Logger
)
