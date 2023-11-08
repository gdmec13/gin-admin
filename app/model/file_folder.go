package model

type FileFolder struct {
	BaseModel
	FileFolderName string
	ParentFolderId int64
	FileStoreId    int64
	Time           string
}
