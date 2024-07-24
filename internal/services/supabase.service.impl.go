package services

import (
	"io"

	"launchpad.net/goamz/s3"
)
type S3ServiceImpl struct {
    client *s3.S3
}



func NewS3ServiceImpl(client *s3.S3) *S3ServiceImpl {
    return &S3ServiceImpl{client: client}
}

// Upload implements UploadService.
func (s *S3ServiceImpl) Upload(bucketName string, fileName string, data io.Reader) string {
		// Upload the file
		// _, err := s.client.UploadFile(bucketName,fileName, data)
		// if err != nil {
		// 	fmt.Errorf("Failed to upload file: %v", err)
		// 	return ""
		// }
		
		// // Get the public URL of the file
		// url := s.client.GetPublicUrl(bucketName,fileName)
		// return url.SignedURL
	return ""
}