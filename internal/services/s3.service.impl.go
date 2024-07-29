package services

import (
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

type StorageServiceImpl struct {
	storage *storage_go.Client
}

func NewStorageServiceImpl(storage *storage_go.Client) StorageService {
	return &StorageServiceImpl{storage: storage}
}

// Upload implements UploadService.
func (s *StorageServiceImpl) Upload(bucketName string, fileName string, bucketKey string) error {

	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()
	
	return nil
}
