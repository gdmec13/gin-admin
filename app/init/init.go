package init

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"simple-cloud-storage/app/global"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("load config file err: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
		if err := v.Unmarshal(&global.APP_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.APP_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.APP_VP = v
}
