package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {


	creds := credentials.NewStaticCredentials("none","none","none")

	config := &aws.Config{
		Endpoint: aws.String("http://0.0.0.0:3003"),
		Region:   aws.String("us-east-1"),
		Credentials:creds,
	}

	sess, err := session.NewSession(config)

	if err != nil {
		panic(err)
	}

	ssmManager := ssm.New(sess)
	query := ssm.GetParameterInput{
		Name:           aws.String("/dev/param1"),
		WithDecryption: aws.Bool(true),
	}

	result, err := ssmManager.GetParameter(&query)

	fmt.Println(*result.Parameter.Value)
}