FROM golang:1.12

WORKDIR /go/src/server

COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on go mod download

ADD . /go/src/server

RUN GO111MODULE=on go build server.go

CMD env=$ENV ./server $PORT
