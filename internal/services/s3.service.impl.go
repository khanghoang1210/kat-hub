package services

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3ServiceImpl struct {
	s3 *s3.Client
}

func NewS3ServiceImpl(s3 *s3.Client) S3Service {
	return &S3ServiceImpl{s3: s3}
}

// Upload implements UploadService.
func (s *S3ServiceImpl) Upload(bucketName string, fileName string, bucketKey string) error {

	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()

	_, err := s.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key: aws.String(bucketKey),
		Body: file,
	})
	if err != nil {
	 return err
	}
	return nil
}
