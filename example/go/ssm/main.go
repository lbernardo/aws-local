package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/lbernardo/aws-local/awslocal"
)

func main() {

	awslocal.SetLocalDev() // Set env AWSLOCAL_DEV=OK

	s,_ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})

	sess := awslocal.GetSessionAWS("http://0.0.0.0:3003", s)

	ssmManager := ssm.New(sess)
	query := ssm.GetParameterInput{
		Name:           aws.String("/dev/param1"),
		WithDecryption: aws.Bool(true),
	}

	result, _ := ssmManager.GetParameter(&query)

	fmt.Println(*result.Parameter.Value)
}