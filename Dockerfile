# Start from Golang v1.14 base image to build the server
FROM golang:1.14.4-alpine3.12 as build

# ENV GO111MODULE=on

# Install git & Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Clone this repo
RUN git clone https://github.com/OhMinsSup/story-server.git /app

# Change workdir
WORKDIR /app

# Copy go mod and sum files 작업 공간에 go.mod 및 go.sum 파일 복사
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 의존성 mod / sum을 변경하지 않으면 캐시됩니다.
RUN go mod download

# Build server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o story-server  .

# Run the server in this container #
FROM alpine:3.12

WORKDIR /app

COPY --from=build /app/.env .
COPY --from=build /app/main .

CMD ["./story-server"]
