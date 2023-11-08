package model

type Share struct {
	BaseModel
	Code     string
	FileId   int64
	Username string
	Hash     string
}
