# go-api

## Main purpose
This project is an example of a quick-start project developed in golang, exposing a rest API

## Technical details
The rest api is exposed using `gin`

The application is deliver through a docker image.

Two docker images: 
- one based on the official `golang` image for the build part and and deployed on a distroless one.
- one based on a distroless image in order to reduce image size.


### debug

Install delve :

`go get -u github.com/go-delve/delve/cmd/dlv`

use this command :

`dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient`

### build project
In order to build project, as usual `bash build.sh`

### Launch tests
In order to launch all tests : `go test ./...`

If you need to launch one specific test : `go test -run path/to/file_test.go`

### Docker commands

run docker

`docker-compose up -d` from the package folder

## How to test it

#### get state
```
curl  http://localhost:8080/info
```

# Generate mocks
In order to generate mock, we use Mockery. 

First install it : 
`go install github.com/vektra/mockery/v2@v2.38.0`

Then generate mock into src folder using the following command
>mockery
