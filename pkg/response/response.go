package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, CommonResponse{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Success(c *gin.Context) {
	result(0, map[string]interface{}{}, "success", c)
}

func SuccessWithMsg(message string, c *gin.Context) {
	result(0, map[string]interface{}{}, message, c)
}

func SuccessDetailed(data interface{}, message string, c *gin.Context) {
	result(0, data, message, c)
}

func FailWithError(message error, errorCode int, c *gin.Context) {
	result(-1, map[string]interface{}{}, fmt.Sprintf("%v", message), c)
}
