package api

import (
	"github.com/gin-gonic/gin"
	"simple-cloud-storage/pkg/response"
)

type FilesStruct struct {
	OpenId string `form:"openId" binding:"required"`
	FId    int    `form:"fId" default:"0"`
}

func Files(c *gin.Context) {
	var R FilesStruct
	if err := c.ShouldBind(&R); err != nil {
		response.FailWithError(err, -1, c)
	} else {
		response.SuccessDetailed(R, "success", c)
	}
}
