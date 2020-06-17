package core

import "io"

type ExecuteLambdaRequest struct {
	Volume      string
	Net         string
	Handler     string
	Runtime     string
	Environment map[string]string
	Body        io.ReadCloser
	Headers     map[string]string
	Parameters  map[string]string
}

type ResultLambdaRequest struct {
	StatusCode        int               `json:"statusCode"`
	Headers           map[string]string `json:"headers,omitempty"`
	MultiValueHeaders string            `json:"multiValueHeaders,omitempty"`
	Body              string            `json:"body"`
}

type ContentRequest struct {
	Body           string            `json:"body"`
	PathParameters map[string]string `json:"pathparameters"`
	Headers        map[string]string `json:"headers"`
}

type Serverless struct {
	Functions map[string]Functions `json:"functions"`
	Provider  Provider             `json:"provider"`
}

type Functions struct {
	Events  []Event `json:"events"`
	Handler string  `json:"handler"`
}

type Event struct {
	HttpEvent HttpEvent `json:"http"`
}

type HttpEvent struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type Provider struct {
	Environment map[string]string `json:"environment"`
	Name        string            `json:"name"`
	Runtime     string            `json:"runtime"`
}
