## Proctor

Proctor is an automation framework. It helps everyone contribute to automation, mange it and use it.

### Introduction

Proctor Engine is the heart of the automation framework.
It takes care of executing jobs and maintaining their configuration.

### Dev environment setup

* Install and setup golang
* Install glide
* Make a directory `src/github.com/gojekfarm` inside your GOPATH
* Clone this repo inside above directory
* Install dependencies using glide
* Install kubectl
* Configure kubectl to point to desired kubernetes cluster. For setting up kubernetes cluster locally, refer [here](https://kubernetes.io/docs/getting-started-guides/minikube/)
* Run a kubectl proxy server on your local machine
* [Configure proctor-engine](#proctor-engine-configuration)
* Running `go build -o proctor-engine` will build binary to run proctor-engine 
* Start service by `./proctor-engine start`
* Run `curl {host-address:port}/ping` for health-check of service

### Running tests

* [Setup dev environment](#dev-environment-setup)
* Run tests: `go test -race -cover $(glide novendor)`

#### Proctor Engine configuration

* Copy `.env.sample` into `.env` file
* Please refer meaning of proctor-engine configuration [here](#proctor-engine-configuration-explanation)
* Modify configuration for dev setup in `.env` file
* Export environment variables configured in `.env` file. Proctor engine gets configuration from environment variables

#### Proctor Engine configuration explanation

* `PROCTOR_APP_PORT` is port on which service will run
* `PROCTOR_LOG_LEVEL` defines log levels of service. Available options are: `debug`,`info`,`warn`,`error`,`fatal`,`panic`
* `PROCTOR_REDIS_ADDRESS` is hostname and port of redis store for jobs configuration and metadata
* `PROCTOR_REDIS_MAX_ACTIVE_CONNECTIONS` defines maximum active connections to redis. Maximum idle connections is half of this config 
* `PROCTOR_LOGS_STREAM_READ_BUFFER_SIZE` and `PROCTOR_LOGS_STREAM_WRITE_BUFFER_SIZE` is the buffer size for websocket connection while streaming logs
* `PROCTOR_KUBE_CONFIG` needs to be set only if service is running outside a kubernetes cluster
  * If unset, service will execute jobs in the same kubernetes cluster where it is run
  * When set to "out-of-cluster", service will fetch kube config based on current-context from `.kube/config` file in home directory
* If a job doesn't reach completion, it is terminated after `PROCTOR_KUBE_JOB_ACTIVE_DEADLINE_SECONDS`
* `PROCTOR_DEFAULT_NAMESPACE` is the namespace under which jobs will be run in kubernetes cluster
* `PROCTOR_KUBE_CLUSTER_HOST_NAME` is address to proxy server for kube cluster. It is used for fetching logs of a pod using http
* Before streaming logs of jobs, `PROCTOR_KUBE_POD_LIST_WAIT_TIME` is the time to wait until jobs and pods are in active/successful/failed state
