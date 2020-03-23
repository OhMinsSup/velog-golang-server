- run:
	@go run main.go

- build-docker:
	@docker build -t story-api:lastest .
	@docker container run -d -p 4000:4000 story-api:lastest
