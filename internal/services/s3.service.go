package services



type S3Service interface {
	Upload(bucketName string, fileName string, bucketKey string) error
}