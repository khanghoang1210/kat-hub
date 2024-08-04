package services

import "io"

type StorageService interface {
	Upload(bucketName string, fileName string, data io.Reader, contentType string) error
	GetPublicUrl(bucketName string, fileName string) string
}