package ssm

import "time"

type Params struct {
	Port string
	Host string
	EnvFile string
}

type Request struct {
	Name           string `json:"Name"`
	WithDecryption bool   `json:"WithDecryption"`
}

type Response struct {
	Parameter Parameter `json:"Parameter"`
}

type Parameter struct {
	ARN              string    `type:"string"`
	LastModifiedDate time.Time `type:"timestamp"`
	Name             string    `min:"1" type:"string"`
	Selector         string    `type:"string"`
	SourceResult     string    `type:"string"`
	Type             string    `type:"string"`
	Value            string    `type:"string"`
	Version          int64     `type:"long"`
}
