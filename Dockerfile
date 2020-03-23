# Run command below to build binary.
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o main main.go
# CGO_ENABLED=0 : cgo를 사용하지 않습니다. Scratch 이미지에는 C 바이너리조차 없기 때문에, 반드시 cgo를 비활성화 후 빌드해야합니다.
# GOOS=linux GOARCH=amd64 : OS와 아키텍쳐 설정입니다.
# -a : 모든(all) 의존 패키지를 cgo를 사용하지 않도록 재빌드합니다.
# -ldflags '-s' : 바이너리를 조금 더 경량화하는 Linker 옵션입니다.

### Builder
FROM golang:1.13-alpine as builder
RUN apk update && apk add git && apk add ca-certificates

WORKDIR /usr/src/app

# 작업 공간에 go.mod 및 go.sum 파일 복사
COPY . .

#의존성 mod / sum을 변경하지 않으면 캐시됩니다.
RUN go mod download

# binary 빌드
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .

# 이미지 크기를 작게
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/src/app/main /main

CMD [ "/main" ]
