package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/lbernardo/aws-local/awslocal"
	"os"
)

func main() {
	// It's developer/test code
	awslocal.SetLocalDev()
	filename := "teste.txt"

	s,_ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})

	// If ENV AWSLOCAL_DEV is "OK" return session endpoint, else return session created
	sess := awslocal.GetSessionAWS("http://0.0.0.0:3002", s)

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

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
