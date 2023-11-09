package model

import (
	"errors"
	"gorm.io/gorm"
	"simple-cloud-storage/app/global"
)

type FileStore struct {
	BaseModel
	UserId      int
	CurrentSize int
	MaxSize     int
}

func CreateFileStore(db *gorm.DB, user User) (int, error) {
	fileStore := FileStore{
		UserId:      user.ID,
		CurrentSize: 0,
		MaxSize:     1048576,
	}
	result := db.Create(&fileStore)
	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected != 1 {
		return 0, errors.New("CreateFileStore fail")
	}
	return fileStore.ID, nil
}

func GetUserFileStore(userId int) (fileStore FileStore) {
	global.APP_DB.Find(&fileStore, "user_id=?", userId)
	return
}

func CapacityIsEnough(fileSize int, fileStoreId int) bool {
	var fileStore FileStore
	global.APP_DB.First(&fileStore, fileStoreId)
	if fileStore.MaxSize-(fileSize/1024) < 0 {
		return false
	}
	return true
}
