package apigateway

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lbernardo/aws-local/internal/adapters/secondary/docker"
	"github.com/lbernardo/aws-local/pkg/core"
)

func StartApiGateway(params ParamsApiGateway) {
	route := mux.NewRouter()

	docker.PullImageDocker(params.Serverless.Provider.Runtime)

	for _, functions := range params.Serverless.Functions {
		path := "/" + functions.Events[0].HttpEvent.Path
		method := functions.Events[0].HttpEvent.Method
		function := functions.Handler
		path = strings.ReplaceAll(path, "//", "/")

		fmt.Printf("http://%v:%v%v [%v]\n\n", params.Host, params.Port, path, strings.ToUpper(method))

		fff := func(w http.ResponseWriter, r *http.Request) {
			parameters := mux.Vars(r)
			headers := map[string]string{}

			for key, _ := range r.Header {
				headers[key] = r.Header.Get(key)
			}

			result, off := docker.ExecuteDockerLambda(core.ExecuteLambdaRequest{
				Volume:      params.Volume,
				Net:         params.Network,
				Handler:     function,
				Runtime:     params.Serverless.Provider.Runtime,
				Headers:     headers,
				Environment: params.Serverless.Provider.Environment,
				Body:        r.Body,
				Parameters:  parameters,
			})
			if result.StatusCode == 0 {
				w.WriteHeader(400)
				fmt.Println(off)
				return
			}

			for key, val := range result.Headers {
				w.Header().Set(key, val)
			}
			w.WriteHeader(result.StatusCode)
			w.Write([]byte(result.Body))
			return
		}
		route.HandleFunc(path, fff).Methods(method)
	}

	fmt.Printf("Start server API Gateway + lambda %v:%v\n", params.Host, params.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", params.Host, params.Port), route))
}
