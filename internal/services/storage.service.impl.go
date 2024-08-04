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
func (s *StorageServiceImpl) Upload(bucketName string, fileName string, data io.Reader, contentType string) error {

	//contentType := "image/jpeg"
	result, err := s.storage.UploadFile(bucketName,fileName,data, storage_go.FileOptions{
		ContentType: &contentType,
	})
	fmt.Println(s.storage.ListBuckets())
	fmt.Println(result)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *StorageServiceImpl) GetPublicUrl(bucketName string, fileName string) string{
	result := s.storage.GetPublicUrl(bucketName, fileName)
	return result.SignedURL
}