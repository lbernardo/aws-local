package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/lbernardo/aws-local/internal/helpers"
	"github.com/lbernardo/aws-local/pkg/core"
)

func PullImageDocker(runtime string) {
	fmt.Println("Prepare image docker")
	imageName := "docker.io/lambci/lambda:" + runtime
	ctx := context.Background()

	fmt.Println(imageName)
	cli, err := client.NewEnvClient()
	if err != nil {
		helpers.PrintError(err)
	}
	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		helpers.PrintError(err)
	}
	io.Copy(os.Stdout, reader)
}

func ReplaceEnvironment(env string) string {
	return strings.ReplaceAll(env, "${opt:stage, self:provider.stage}", "dev")
}

func ExecuteDockerLambda(content core.ExecuteLambdaRequest) (core.ResultLambdaRequest, string) {
	var result core.ResultLambdaRequest
	var output bytes.Buffer
	var contentRequest core.ContentRequest
	bodyStr := ""
	if content.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(content.Body)
		bodyStr = buf.String()
	}
	var strEnv []string

	imageName := "lambci/lambda:" + content.Runtime

	for n, env := range content.Environment {
		strEnv = append(strEnv, n+"="+ReplaceEnvironment(env))
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		helpers.PrintError(err)
	}

	bodyStr = strings.ReplaceAll(bodyStr, "\t", "")
	bodyStr = strings.ReplaceAll(bodyStr, "\n", "")

	contentRequest.Body = bodyStr
	contentRequest.PathParameters = content.Parameters
	contentRequest.Headers = content.Headers

	jsonRequest, _ := json.Marshal(contentRequest)

	var executeCommand []string
	executeCommand = append(executeCommand, content.Handler)
	executeCommand = append(executeCommand, string(jsonRequest))

	// Network config
	networkingConfig := &network.NetworkingConfig{}

	if content.Net != "" {
		networkingConfig.EndpointsConfig = map[string]*network.EndpointSettings{
			"net": &network.EndpointSettings{
				NetworkID: content.Net,
			},
		}
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   executeCommand,
		Env:   strEnv,
	}, &container.HostConfig{
		Binds: []string{content.Volume + ":/var/task"},
	}, networkingConfig, "")
	if err != nil {
		helpers.PrintError(err)
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})

	if err != nil {
		helpers.PrintError(err)
	}

	cli.ContainerWait(ctx, resp.ID)

	reader, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		helpers.PrintError(err)
	}

	stdcopy.StdCopy(&output, os.Stderr, reader)

	str := output.String()
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		fmt.Println(err)
	}

	cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})

	return result, str
}
