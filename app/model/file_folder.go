package model

type FileFolder struct {
	BaseModel
	FileFolderName string
	ParentFolderId int
	FileStoreId    int
	Time           string
}
