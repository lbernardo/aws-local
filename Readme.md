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

###API Gateway
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

**values.ssm**
```
/dev/param1=Hello1
/prod/param2=Hello2
param3=true
param4=1
```