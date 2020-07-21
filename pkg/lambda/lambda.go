package lambda

import (
	"io/ioutil"
	"strings"

	"github.com/lbernardo/aws-local/internal/adapters/secondary/docker"
	"github.com/lbernardo/aws-local/pkg/core"
)

func StartLambda(p ParamsLambda) {

	rBody := ioutil.NopCloser(strings.NewReader(p.Body))

	docker.PullImageDocker(p.Runtime)
	docker.ExecuteDockerLambda(core.ExecuteLambdaRequest{
		Volume:      p.Volume,
		Handler:     p.Bin,
		Runtime:     p.Runtime,
		Environment: p.Env,
		Body:        rBody,
	})
}
