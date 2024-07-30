package services

import "io"

type StorageService interface {
	Upload(bucketName string, fileName io.Reader) error
}