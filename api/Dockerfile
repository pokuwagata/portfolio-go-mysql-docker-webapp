FROM golang:1.12

ADD . /go/src/server
WORKDIR /go/src/server

RUN GO111MODULE=on go build server.go

CMD env=$ENV ./server $PORT
