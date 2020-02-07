package apigateway

import "github.com/lbernardo/aws-local/pkg/core"

type ParamsApiGateway struct {
	Serverless     core.Serverless
	Volume         string
	Host           string
	Port           string
	Network        string
}
