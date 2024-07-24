package services

import "io"


type S3Service interface {
	Upload(bucketName string, fileName string, data io.Reader) string
}