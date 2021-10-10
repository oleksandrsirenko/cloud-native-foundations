
# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /go/src/app

ADD . .

RUN go mod init

RUN go build -o helloworld

EXPOSE 8080

CMD ["./helloworld"]