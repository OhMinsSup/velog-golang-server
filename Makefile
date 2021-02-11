#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=story-server
export LDFLAGS="-w -s"
export APP_ENV="production"

- db-sync:
	go generate ./ent

- run:
	APP_ENV="development" go run main.go

- start:
	APP_ENV="production" go run main.go

- build:
	APP_ENV="production" go build -race  .

- build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .


