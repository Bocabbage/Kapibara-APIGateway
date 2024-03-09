FROM golang:1.21

WORKDIR /app
ADD . /app

RUN go env -w  GOPROXY=https://goproxy.cn
RUN go build -o ./kapibara-apigateway_exe ./cmd/apiserver/main.go
