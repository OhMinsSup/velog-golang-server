#gin -p 4001 -a 4000 run main.go
- run:
	@go run main.go

- build-docker:
	@docker build -t story-api:lastest .
	@docker container run -d -p 4000:4000 story-api:lastest
