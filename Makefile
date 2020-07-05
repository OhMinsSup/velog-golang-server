#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=story-server
export LDFLAGS="-w -s"

- run:
	@go run main.go

- build:
	@go build -race  .

- build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

- build-start:
	@docker-compose up

- build-down:
	@docker-compose down
