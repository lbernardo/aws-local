package ssm

import (
	"encoding/json"
	"fmt"
	"github.com/lbernardo/aws-local/internal/adapters/secondary/env"
	"io/ioutil"
	"net/http"
	"os"
)

type SSMLocal struct {
	params Params
}

func NewSSMLocal(params Params) *SSMLocal {
	return &SSMLocal{
		params: params,
	}
}

func (s *SSMLocal) getVars() map[string]string {
	fileEnv, err := os.Open(s.params.EnvFile)
	if err != nil {
		panic(err)
	}
	defer fileEnv.Close()
	envs, err := env.Parse(fileEnv)
	if err != nil {
		panic(err)
	}
	return envs
}

func (s *SSMLocal) Start() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var body Request

		if r.Method == "POST" && r.RequestURI == "/" {
			// Values
			values := s.getVars()
			content, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(content, &body)

			if value := values[body.Name] ; value != "" {
				response, _ := json.Marshal(Response{
					Parameter: Parameter{
						Name:  body.Name,
						Value: values[body.Name],
					},
				})
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-type", "application/json")
				w.Write(response)
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ParameterNotFound"))

		}
	})
	fmt.Printf("Start SSM server %v:%v\n", s.params.Host, s.params.Port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", s.params.Host, s.params.Port), nil)
}
