package upload

import (
	"mime/multipart"

	"bookstore/global"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

//OSS接口
//OSS的实例化方法

func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		return &Local{}
	default:
		return &Local{}
	}
}
