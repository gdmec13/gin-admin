package request

type FilesStruct struct {
	OpenId string `form:"openId" binding:"required"`
	FId    int    `form:"fId" default:"0"`
}
