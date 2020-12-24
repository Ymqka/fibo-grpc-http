FROM golang:1.15.6-buster

WORKDIR /src/github.com/Ymqka/fibo-grpc-http
ADD . /src/github.com/Ymqka/fibo-grpc-http

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o bin/serve_fibo_http_grpc

CMD ["./bin/serve_fibo_http_grpc"]
