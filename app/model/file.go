package model

type MyFile struct {
	BaseModel
	FileName       string //文件名
	FileHash       string //文件哈希值
	FileStoreId    int64  //文件仓库id
	FilePath       string //文件存储路径
	DownloadNum    int64  //下载次数
	UploadTime     string //上传时间
	ParentFolderId int64  //父文件夹id
	Size           int64  //文件大小
	SizeStr        string //文件大小单位
	Type           uint8  //文件类型
	Postfix        string //文件后缀
}
