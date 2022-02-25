package dmp

// PopulationUploadFileRequest 人群包文件上传接口
type PopulationUploadFileRequest struct {
	PopulationUploadRequest
}

// Url implement UploadRequest interface
func (r PopulationUploadFileRequest) Url() string {
	return "gw/dmp/v2/dmp/population/file/upload"
}
