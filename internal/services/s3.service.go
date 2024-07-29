package services



type StorageService interface {
	Upload(bucketName string, fileName string, bucketKey string) error
}