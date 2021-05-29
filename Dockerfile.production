FROM golang:1.16.4

ADD . /go/src/github.com/gvu0110/bookstore_users-api

RUN go install github.com/gvu0110/bookstore_users-api@latest

ENTRYPOINT /go/bin/bookstore_users-api

EXPOSE 8081