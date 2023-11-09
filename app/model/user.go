package model

import (
	"errors"
	"gorm.io/gorm"
	"simple-cloud-storage/app/global"
	"time"
)

type User struct {
	BaseModel
	OpenId       string
	FileStoreId  int
	UserName     string
	RegisterTime time.Time
	ImagePath    string
}

func CreateUserAndFileStore(openId, username, image string) error {
	tx := global.APP_DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, e := CreateUser(tx, openId, username, image)
	if e != nil {
		return e
	}

	fileStoreId, err2 := CreateFileStore(tx, user)
	if err2 != nil {
		return err2
	}

	user.FileStoreId = fileStoreId
	tx.Save(&user)

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func GetUserInfo(openId interface{}) (user User) {
	global.APP_DB.Find(&user, "open_id=?", openId)
	return
}

func QueryUserExists(openId string) bool {
	var user User
	global.APP_DB.Find(&user, "open_id = ?", openId)
	if user.ID == 0 {
		return false
	}
	return true
}

func CreateUser(db *gorm.DB, openId, username, image string) (user User, err error) {
	user = User{
		OpenId:       openId,
		FileStoreId:  0,
		UserName:     username,
		RegisterTime: time.Now(),
		ImagePath:    image,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected != 1 {
		return user, errors.New("CreateUser fail")
	}

	return user, nil
}
