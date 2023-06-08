package example

import (
	"errors"
	"mime/multipart"
	"strings"

	"bookstore/global"
	"bookstore/model/common/request"
	"bookstore/model/example"
	"bookstore/utils/upload"
)

//创建文件上传记录

func (e *FileUploadAndDownloadService) Upload(file example.ExaFileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

// 删除文件切片记录
func (e *FileUploadAndDownloadService) FindFile(id int) (error, example.ExaFileUploadAndDownload) {
	var file example.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return err, file
}

// 删除文件记录
func (e *FileUploadAndDownloadService) DeleteFile(file example.ExaFileUploadAndDownload) (err error) {
	var fileFromDb example.ExaFileUploadAndDownload
	err, fileFromDb = e.FindFile(file.ID)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// 分页获取数据
func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	db := global.GVA_DB.Model(&example.ExaFileUploadAndDownload{})
	var fileLists []example.ExaFileUploadAndDownload
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

//根据配置文件判断是文件上传到本地或者七牛云

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (err error, file example.ExaFileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := example.ExaFileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return err, f
	}
	return
}
