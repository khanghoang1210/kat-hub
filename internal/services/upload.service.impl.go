package services

import "io"

type UploadServiceImpl struct {
}



func NewUploadServiceImpl() UploadService {
	return &UploadServiceImpl{}
}

// Upload implements UploadService.
func (u *UploadServiceImpl) Upload(bucketName string, fileName string, data io.Reader) string {
	panic("unimplemented")
}