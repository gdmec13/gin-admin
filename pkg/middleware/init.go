package middleware

import "github.com/gin-gonic/gin"

func SetupCommonMiddleware(r *gin.Engine) {
	r.Use(Cors())
}
