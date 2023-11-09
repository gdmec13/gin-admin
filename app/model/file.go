package model

import (
	"path"
	"simple-cloud-storage/app/global"
	"simple-cloud-storage/pkg/util"
	"strconv"
	"strings"
	"time"
)

const (
	FileTypeDoc   = 1
	FileTypeImage = 2
	FileTypeVideo = 3
	FileTypeMusic = 4
	FileTypeOther = 5
)

type MyFile struct {
	BaseModel
	FileName       string //文件名
	FileHash       string //文件哈希值
	FileStoreId    int    //文件仓库id
	FilePath       string //文件存储路径
	DownloadNum    int    //下载次数
	UploadTime     string //上传时间
	ParentFolderId int    //父文件夹id
	Size           int64  //文件大小
	SizeStr        string //文件大小单位
	Type           int    //文件类型
	Postfix        string //文件后缀
}

// 创建文件对象
func createMyFile(filePrefix, fileHash string, fileStoreId, fid int, fileSize int64, sizeStr, fileSuffix string) *MyFile {
	return &MyFile{
		FileName:       filePrefix,
		FileHash:       fileHash,
		FileStoreId:    fileStoreId,
		FilePath:       "",
		DownloadNum:    0,
		UploadTime:     time.Now().Format("2006-01-02 15:04:05"),
		ParentFolderId: fid,
		Size:           fileSize / 1024,
		SizeStr:        sizeStr,
		Type:           util.GetFileTypeInt(fileSuffix),
		Postfix:        strings.ToLower(fileSuffix),
	}
}

// CreateFile 创建一个新的文件记录
// 参数:
//   - fileName: 文件名
//   - fileHash: 文件哈希
//   - fileSize: 文件大小字节
//   - parentFolderID: 父文件夹ID
//   - fileStoreID: 文件存储ID
// 返回:
//   - 文件记录对象
//   - 错误信息
func CreateFile(filename, fileHash string, fileSize int64, fId string, fileStoreId int) (*MyFile, error) {
	var sizeStr string
	fileSuffix := path.Ext(filename)
	filePrefix := filename[0 : len(filename)-len(fileSuffix)]
	fid, _ := strconv.Atoi(fId)
	sizeStr = util.CalculateSizeStr(fileSize)
	myFile := createMyFile(filePrefix, fileHash, fileStoreId, fid, fileSize, sizeStr, fileSuffix)
	if err := global.APP_DB.Create(myFile).Error; err != nil {
		return nil, err
	}

	return myFile, nil
}

func GetUserFile(parentId string, storeId int) (files []MyFile) {
	global.APP_DB.Find(&files, "file_store_id = ? and parent_folder_id = ?", storeId, parentId)
	return
}

func SubtractSize(size int64, fileStoreId int) error {
	var fileStore FileStore
	if err := global.APP_DB.First(&fileStore, fileStoreId).Error; err != nil {
		return err
	}

	fileStore.CurrentSize += int(size / 1024)
	fileStore.MaxSize -= int(size / 1024)

	if err := global.APP_DB.Save(&fileStore).Error; err != nil {
		return err
	}

	return nil
}

func GetUserFileCount(fileStoreId int) (fileCount int64) {
	var file []MyFile
	global.APP_DB.Find(&file, "file_store_id = ?", fileStoreId).Count(&fileCount)
	return
}

func getCountByType(fileStoreID int, fileType int) (int64, error) {
	var count int64

	if err := global.APP_DB.Where("file_store_id = ? and type = ?", fileStoreID, fileType).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetFileDetailUse 获取文件存储的文件类型分布
// 返回各类型文件的数量映射
func GetFileDetailUse(fileStoreID int) (map[string]int64, error) {

	docCount, err := getCountByType(fileStoreID, FileTypeDoc)
	if err != nil {
		return nil, err
	}

	imgCount, err := getCountByType(fileStoreID, FileTypeImage)
	if err != nil {
		return nil, err
	}

	videoCount, err := getCountByType(fileStoreID, FileTypeVideo)
	if err != nil {
		return nil, err
	}

	musicCount, err := getCountByType(fileStoreID, FileTypeMusic)
	if err != nil {
		return nil, err
	}

	otherCount, err := getCountByType(fileStoreID, FileTypeOther)
	if err != nil {
		return nil, err
	}

	return map[string]int64{
		"docCount":   docCount,
		"imgCount":   imgCount,
		"videoCount": videoCount,
		"musicCount": musicCount,
		"otherCount": otherCount,
	}, nil
}

func GetTypeFile(fileType, fileStoreId int) (files []MyFile) {
	global.APP_DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, fileType)
	return
}

func CurrFileExists(fId, filename string) bool {
	var file MyFile
	//获取文件后缀
	fileSuffix := strings.ToLower(path.Ext(filename))
	//获取文件名
	filePrefix := filename[0 : len(filename)-len(fileSuffix)]

	global.APP_DB.Find(&file, "parent_folder_id = ? and file_name = ? and postfix = ?", fId, filePrefix, fileSuffix)

	if file.Size > 0 {
		return false
	}
	return true
}

func FileOssExists(fileHash string) bool {
	var file MyFile
	global.APP_DB.Find(&file, "file_hash = ?", fileHash)
	if file.FileHash != "" {
		return false
	}
	return true
}

func GetFileInfo(fId string) (file MyFile) {
	global.APP_DB.First(&file, fId)
	return
}

func DownloadNumAdd(fId string) {
	var file MyFile
	global.APP_DB.First(&file, fId)
	file.DownloadNum = file.DownloadNum + 1
	global.APP_DB.Save(&file)
}

func DeleteUserFile(fId, folderId string, storeId int) {
	global.APP_DB.Where("id = ? and file_store_id = ? and parent_folder_id = ?", fId, storeId, folderId).Delete(MyFile{})
}
