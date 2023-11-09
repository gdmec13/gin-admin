package util

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strconv"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GenerateFile(filename, content string) error {
	// 创建文件并打开
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入内容
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func GetFileTypeInt(filePrefix string) int {
	filePrefix = strings.ToLower(filePrefix)
	if filePrefix == ".doc" || filePrefix == ".docx" || filePrefix == ".txt" || filePrefix == ".pdf" {
		return 1
	}
	if filePrefix == ".jpg" || filePrefix == ".png" || filePrefix == ".gif" || filePrefix == ".jpeg" {
		return 2
	}
	if filePrefix == ".mp4" || filePrefix == ".avi" || filePrefix == ".mov" || filePrefix == ".rmvb" || filePrefix == ".rm" {
		return 3
	}
	if filePrefix == ".mp3" || filePrefix == ".cda" || filePrefix == ".wav" || filePrefix == ".wma" || filePrefix == ".ogg" {
		return 4
	}

	return 5
}

// 计算文件大小字符串
func CalculateSizeStr(fileSize int64) string {
	var sizeStr string
	if fileSize < 1048576 {
		sizeStr = strconv.FormatInt(fileSize/1024, 10) + "KB"
	} else {
		sizeStr = strconv.FormatInt(fileSize/102400, 10) + "MB"
	}
	return sizeStr
}
