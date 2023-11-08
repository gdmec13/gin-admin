package model

type FileStore struct {
	BaseModel
	UserId      int64
	CurrentSize int64
	MaxSize     int64
}
