FROM golang:1.21

WORKDIR /app
ADD . /app

EXPOSE 5678

RUN go build -o ./kapibara-apigateway_exe ./cmd/apiserver/main.go

ENTRYPOINT ["./kapibara-apigateway_exe", "-tls"]
# ENTRYPOINT [ "tail", "-f", "/dev/null" ]