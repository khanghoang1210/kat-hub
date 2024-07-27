package initialize

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

// CreateSession creates a new AWS session
func CreateSession() {
	var err error
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("default"),
		config.WithRegion("ap-southeast-1"),
	)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(cfg)

	S3Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String("https://urevhrzshuvxvvkuduco.supabase.co/storage/v1/s3")
		o.UsePathStyle = true
	})

}
