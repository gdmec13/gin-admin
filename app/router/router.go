package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"simple-cloud-storage/app/api"
	"simple-cloud-storage/app/global"
	OpInit "simple-cloud-storage/app/init"
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

	// 初始化数据库
	OpInit.SetupMysql()
	defer global.CloseDb()

	// 生成数据表
	OpInit.GenerateTable()

	// 设置gin的运行模式
	gin.SetMode(global.APP_CONFIG.System.RunModel)
	global.APP_LOG.Debug("GIN Set Model:", global.APP_CONFIG.System.RunModel)

	httpPort := fmt.Sprintf(":%v", global.APP_CONFIG.System.Port)
	r := gin.Default()
	middleware.SetupCommonMiddleware(r)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Web!")
	})

	// 注册其它路由
	registerRouter(r)

	if err := r.Run(httpPort); err != nil {
		log.Fatal("run server fail...")
	}

}

func registerRouter(r *gin.Engine) {
	r.GET("/files", api.Files)
	r.GET("/login", api.Login)
}
