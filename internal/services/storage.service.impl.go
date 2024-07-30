package services

import (
	"fmt"
	"io"

	storage_go "github.com/supabase-community/storage-go"
)

type StorageServiceImpl struct {
	storage *storage_go.Client
}

func NewStorageServiceImpl(storage *storage_go.Client) StorageService {
	return &StorageServiceImpl{storage: storage}
}

// Upload implements UploadService.
func (s *StorageServiceImpl) Upload(bucketName string, file io.Reader) error {

	result, err := s.storage.UploadFile(bucketName,"choigem1.jpg",file)
	fmt.Println(s.storage.ListBuckets())
	fmt.Println(result)
	if err != nil {
		return err
	}
	
	return nil
}
