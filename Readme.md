# aws-local
Execute and test local services AWS


```bash
$ awslocal 
Start modules awslocal

Usage:
  awslocal [command]

Available Commands:
  api-gateway API Gateway with Lambda
  help        Help about any command
  s3          S3
  ssm         Security Secret Manager (SSM)

Flags:
  -h, --help   help for awslocal
```

## Modules

- api-gateway
- S3
- SSM

---

### API Gateway
Use serverless framework for test API Gateway + Lambda

```
Usage:
  awslocal api-gateway [flags]

Flags:
      --env string       File for using environment variables other than serverless. Can replace serverless variables
  -h, --help             help for api-gateway
      --host string      host usage [default 0.0.0.0] (default "0.0.0.0")
      --network string   Set network name usage
      --port string      port usage [default 3000] (default "3000")
  -v, --volume string    Volume project (Ep: --volume $PWD) (required)
      --yaml string      Serverless file yaml (default "serverless.yml")
```

#### Start service
```
awslocal api-gateway  --volume $PWD 
```

> **--volume** is required. **Volume** is source path of project

> If you want to connect to the Dynamodb database on a docker, you need to pass **--network** for communication between networks

> Use the --env file to create environment variables that can be used by lambda



### S3
Use S3 storage local for test and development

```
Usage:
  awslocal s3 [flags]

Flags:
  -h, --help             help for s3
      --host string      host usage [default 0.0.0.0] (default "0.0.0.0")
      --network string  Set network name usage
      --port string     port usage [default 3002] (default "3002")
  -v, --volume string   Volume for storage S3 [required]
 ```

#### Start service
```
awslocal s3  --volume /home/dev/mystorage 
```

**Example code**
```go
package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/lbernardo/aws-local/awslocal"
	"os"
)

func main() {
	// It's developer/test code
	awslocal.SetLocalDev() // Set env AWSLOCAL_DEV=OK
	filename := "teste.txt"

	s,_ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})

	// If ENV AWSLOCAL_DEV is "OK" return session endpoint, else return session created
	sess := awslocal.GetSessionAWS("http://0.0.0.0:3002", s)

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err  := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	data, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("/xxxy/teste.txt"),
		Body:   f,
	})
	fmt.Println(data)
	if err != nil {
		fmt.Errorf("failed to upload file, %v", err)
	}
}
```

### SSM
Use Secret Security Manager for devs and tests

```
Security Secret Manager (SSM)

Usage:
  awslocal ssm [flags]

Flags:
  -h, --help            help for ssm
      --host string     host usage [default 0.0.0.0] (default "0.0.0.0")
      --port string     port usage [default 3003] (default "3003")
  -v, --values string   File using for SSM values [required]

```


#### Start service
```bash
awslocal ssm --values values.ssm
```

**values.ssm** (example)
```
/dev/param1=Hello1
/prod/param2=Hello2
param3=true
param4=1
```


**Example code**
```go
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

    // If ENV AWSLOCAL_DEV is "OK" return session endpoint, else return session created
    sess := awslocal.GetSessionAWS("http://0.0.0.0:3003", s)

	ssmManager := ssm.New(sess)
	query := ssm.GetParameterInput{
		Name:           aws.String("/dev/param1"),
		WithDecryption: aws.Bool(true),
	}

	result, _ := ssmManager.GetParameter(&query)

	fmt.Println(*result.Parameter.Value)
}
```

## Create new modules


##### Create command

*pkg/mymodule/mymodule.go*
```go
package mymodule

import (
	"fmt"
	"net/http"
)

type MyModule struct {
	params Params
}

func NewMyModule(params Params) *MyModule {
	return &MyModule{
		params: params,
	}
}


func (m *MyModule) StartModule() {
	fmt.Printf("Start MyModule server %v:%v\n",m.params.Host, m.params.Port)
	http.ListenAndServe(fmt.Sprintf("%v:%v",m.params.Host, m.params.Port),nil)
}
```

*pkg/mymodule/model.go*
```go
package mymodule

type Params struct {
	Host string
	Port string
}
```

*cmd/mymodule.go*
```go
package cmd

import (
	"github.com/lbernardo/aws-local/pkg/mymodule"
	"github.com/spf13/cobra"
)

var myModuleCommand = &cobra.Command{
	Use:   "mymodule",
	Short: "My module describe",
	Run:   ExecuteMyModuleCommand,
}
func init() {
	myModuleCommand.PersistentFlags().StringVar(&portMyModule, "port", "3004", "port usage [default 3002]")
	myModuleCommand.PersistentFlags().StringVar(&hostMyModule, "host", "0.0.0.0", "host usage [default 0.0.0.0]")

	rootCmd.AddCommand(myModuleCommand)
}

var portMyModule string
var hostMyModule string


func ExecuteMyModuleCommand(cmd *cobra.Command, args []string) {

	mymodule.NewMyModule(mymodule.Params{
		Host: hostMyModule,
		Port: portMyModule,
	}).StartModule()

}
```

*Execute mymodule*
```
awslocal mymodule
```