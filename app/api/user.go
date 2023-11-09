package api

import (
	"github.com/gin-gonic/gin"
	"simple-cloud-storage/app/api/request"
	"simple-cloud-storage/app/model"
	"simple-cloud-storage/pkg/response"
)

func Login(c *gin.Context) {
	var R request.FilesStruct
	if err := c.ShouldBind(&R); err != nil {
		response.FailWithError(err, -1, c)
	} else {
		err := model.CreateUserAndFileStore(R.OpenId, "test", "test")
		if err != nil {
			response.FailWithError(err, -1, c)
			return
		}
		response.SuccessDetailed(R, "success", c)
	}
}
