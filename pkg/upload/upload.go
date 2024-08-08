package upload

import (
	"mime/multipart"

	"go.uber.org/zap"
)

//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@interface_name: OSS
//@description: OSS接口
type OSS interface {
	UploadFile(file *multipart.FileHeader, localPath string, log *zap.Logger) (string, string, error)
	DeleteFile(key string, localPath string, log *zap.Logger) error
}

//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: NewOss
//@description: OSS接口
//@description: OSS的实例化方法
//@return: OSS
func NewOss(ossType string) OSS {
	switch ossType {
	case "local":
		return &Local{}
	// case "qiniu":
	// 	return &Qiniu{}
	// case "tencent-cos":
	// 	return &TencentCOS{}
	// case "aliyun-oss":
	// 	return &AliyunOSS{}
	default:
		return &Local{}
	}
}
