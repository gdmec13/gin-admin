package model

import (
	"gorm.io/gorm"
	"simple-cloud-storage/app/global"
	"strconv"
	"time"
)

type FileFolder struct {
	BaseModel
	FileFolderName string
	ParentFolderId int
	FileStoreId    int
	Time           string
}

func newFolder(folderName string, parentIdInt, fileStoreId int) *FileFolder {
	return &FileFolder{
		FileFolderName: folderName,
		ParentFolderId: parentIdInt,
		FileStoreId:    fileStoreId,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
	}
}

func CreateFolder(folderName string, parentId string, fileStoreId int) (*FileFolder, error) {
	parentIdInt, err := strconv.Atoi(parentId)
	if err != nil {
		return nil, err
	}
	newFolder := newFolder(folderName, parentIdInt, fileStoreId)
	if err := global.APP_DB.Create(newFolder).Error; err != nil {
		return nil, err
	}

	return newFolder, nil
}

func GetParentFolder(fId string) (fileFolder FileFolder) {
	global.APP_DB.Find(&fileFolder, "id = ?", fId)
	return
}

func GetFileFolder(parentId string, fileStoreId int) (fileFolders []FileFolder) {
	global.APP_DB.Order("time desc").Find(&fileFolders, "parent_folder_id = ? and file_store_id = ?", parentId, fileStoreId)
	return
}

func GetCurrentFolder(fId string) (fileFolder FileFolder) {
	global.APP_DB.Find(&fileFolder, "id = ?", fId)
	return
}

func GetCurrentAllParent(folder FileFolder, folders []FileFolder) []FileFolder {
	var parentFolder FileFolder
	if folder.ParentFolderId != 0 {
		global.APP_DB.Find(&parentFolder, "id = ?", folder.ParentFolderId)
		folders = append(folders, parentFolder)
		//递归查找当前所有父级
		return GetCurrentAllParent(parentFolder, folders)
	}

	//反转切片
	for i, j := 0, len(folders)-1; i < j; i, j = i+1, j-1 {
		folders[i], folders[j] = folders[j], folders[i]
	}

	return folders
}

func GetUserFileFolderCount(fileStoreId int) (fileFolderCount int64) {
	var fileFolder []FileFolder
	global.APP_DB.Find(&fileFolder, "file_store_id = ?", fileStoreId).Count(&fileFolderCount)
	return
}

func deleteFileFolder(tx *gorm.DB, folderID int) error {

	// 删除当前文件夹
	if err := tx.Where("id = ?", folderID).Delete(&FileFolder{}).Error; err != nil {
		return err
	}

	// 删除当前文件夹中的文件
	if err := tx.Where("parent_folder_id = ?", folderID).Delete(&MyFile{}).Error; err != nil {
		return err
	}

	// 查找当前文件夹的子文件夹
	var childFolders []FileFolder
	if err := tx.Where("parent_folder_id = ?", folderID).Find(&childFolders).Error; err != nil {
		return err
	}

	// 递归删除每个子文件夹
	for _, child := range childFolders {
		if err := deleteFileFolder(tx, child.ID); err != nil {
			return err
		}
	}

	return nil
}

func DeleteFileFolder(folderID string) error {
	tx := global.APP_DB.Begin()
	defer tx.Rollback()

	if err := tx.Where("id = ?", folderID).Delete(&FileFolder{}).Error; err != nil {
		return err
	}

	if err := tx.Where("parent_folder_id = ?", folderID).Delete(&MyFile{}).Error; err != nil {
		return err
	}

	// 查找子文件夹
	var childFolders []FileFolder
	if err := tx.Where("parent_folder_id = ?", folderID).Find(&childFolders).Error; err != nil {
		return err
	}

	// 递归删除子文件夹
	for _, child := range childFolders {
		if err := deleteFileFolder(tx, child.ID); err != nil {
			return err
		}
	}

	return tx.Commit().Error
}

func UpdateFolderName(fId, fName string) {
	var fileFolder FileFolder
	global.APP_DB.Model(&fileFolder).Where("id = ?", fId).Update("file_folder_name", fName)
}
