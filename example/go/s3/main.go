package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	filename := "teste.txt"

	creds := credentials.NewStaticCredentials("none","none","none")

	config := &aws.Config{
		Endpoint: aws.String("http://0.0.0.0:3002"),
		Region:   aws.String("us-east-1"),
		Credentials:creds,
	}

	sess, err := session.NewSession(config)

	iter := s3manager.NewDeleteListIterator(sess, &s3.ListObjectsInput{
		Bucket: aws.String("myBucket"),

	})

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	d := s3manager.NewBatchDelete(sess)
	d.Delete(aws.Context(),)

	f, err  := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	data, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("/xxxy/teste.txt"),
		Body:   f,
	})
	fmt.Println(data)
	if err != nil {
		fmt.Errorf("failed to upload file, %v", err)
	}
}
