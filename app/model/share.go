package model

import (
	"simple-cloud-storage/app/global"
	"simple-cloud-storage/pkg/util"
	"time"
)

type Share struct {
	BaseModel
	Code     string
	FileId   int
	Username string
	Hash     string
}

func CreateShare(code, username string, fId int) string {
	share := Share{
		Code:     code,
		FileId:   fId,
		Username: username,
		Hash:     util.Md5(code + string(time.Now().Unix())),
	}
	global.APP_DB.Create(&share)
	return share.Hash
}

func GetShareInfo(f string) (share Share) {
	global.APP_DB.Find(&share, "hash=?", f)
	return
}

func VerifyShareCode(fId, code string) bool {
	var share Share
	global.APP_DB.Find(&share, "file_id=? and code=?", fId, code)
	if share.ID == 0 {
		return false
	}
	return true
}
