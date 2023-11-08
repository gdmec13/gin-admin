package model

import "time"

type User struct {
	BaseModel
	OpenId       string
	FileStoreId  int64
	UserName     string
	RegisterTime time.Time
	ImagePath    string
}
