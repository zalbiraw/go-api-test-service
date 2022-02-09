FROM        golang:1.17.6-alpine

MAINTAINER  Zaid Albirawi

WORKDIR     /go/src/github.com/zalbiraw/go-api-test-service

COPY        go.mod      ./go.mod
COPY        users/      ./users
COPY        posts/      ./posts
COPY        comments/   ./comments

RUN         go mod tidy
RUN         go mod vendor
