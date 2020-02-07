package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	filename := "teste.txt"

	config := &aws.Config{
		Endpoint: aws.String("http://127.0.0.1:3002"),
		Region:   aws.String("us-east-1"),
	}

	sess, err := session.NewSession(config)

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err  := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("myString"),
		Body:   f,
	})
	if err != nil {
		fmt.Errorf("failed to upload file, %v", err)
	}
}
