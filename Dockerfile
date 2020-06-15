# Start from golang base image
FROM golang:1.13-alpine as builder

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add git && apk add ca-certificates
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files 작업 공간에 go.mod 및 go.sum 파일 복사
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 의존성 mod / sum을 변경하지 않으면 캐시됩니다.
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app binary 빌드
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch 이미지 크기를 작게
# FROM scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /usr/src/app/main /main

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/.env.dev .
COPY --from=builder /app/.env.prod .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD [ "./main" ]

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o main main.go
# CGO_ENABLED=0 : cgo를 사용하지 않습니다. Scratch 이미지에는 C 바이너리조차 없기 때문에, 반드시 cgo를 비활성화 후 빌드해야합니다.
# GOOS=linux GOARCH=amd64 : OS와 아키텍쳐 설정입니다.
# -a : 모든(all) 의존 패키지를 cgo를 사용하지 않도록 재빌드합니다.
# -ldflags '-s' : 바이너리를 조금 더 경량화하는 Linker 옵션입니다.
