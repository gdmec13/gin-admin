package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"simple-cloud-storage/app/global"
	_ "simple-cloud-storage/app/init"
	OpLog "simple-cloud-storage/pkg/log"
	"simple-cloud-storage/pkg/middleware"
)

func init() {
	InitRouter()
}

func InitRouter() {
	// 注册日志
	OpLog.InitLogger()
	global.APP_LOG.Debug("logger init success")

	gin.SetMode(global.APP_CONFIG.System.RunModel)
	global.APP_LOG.Debug("GIN Set Model:", global.APP_CONFIG.System.RunModel)

	httpPort := fmt.Sprintf(":%v", global.APP_CONFIG.System.Port)
	r := gin.Default()
	middleware.SetupCommonMiddleware(r)

	r.GET("/", func(c *gin.Context) {
		fmt.Println("welcome")
	})

	if err := r.Run(httpPort); err != nil {
		log.Fatal("run server fail...")
	}

}
