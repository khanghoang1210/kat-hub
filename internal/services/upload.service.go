package services

import "io"

type UploadService interface {
	Upload(bucketName string, fileName string, data io.Reader) string
}