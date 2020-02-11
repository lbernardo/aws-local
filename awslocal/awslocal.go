package awslocal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

const AWSLOCAL_DEV = "AWSLOCAL_DEV"

func IsLocalDev() bool {
	return os.Getenv(AWSLOCAL_DEV) == "TRUE"
}

func SetLocalDev() {
	os.Setenv(AWSLOCAL_DEV, "TRUE")
}

func DelLocalDev() {
	os.Unsetenv(AWSLOCAL_DEV)
}

func GetSessionAWS(endpoint string, defaultSession *session.Session) *session.Session {
	if !IsLocalDev() {
		return defaultSession
	}
	creds := credentials.NewStaticCredentials("awslocal","awslocal","awslocal")

	config := &aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String("us-east-1"),
		Credentials:creds,
	}

	sess, _ := session.NewSession(config)
	return sess
}
