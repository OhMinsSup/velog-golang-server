#gin -p 4001 -a 4000 run main.go
- run:
	@go run main.go

- build-start:
	@docker-compose up

- build-down:
	@docker-compose down
