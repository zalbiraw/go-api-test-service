FROM        golang:1.17.6-alpine

MAINTAINER  Zaid Albirawi

ARG         arch

WORKDIR     /go/src/github.com/zalbiraw/go-api-test-service

COPY        go.mod      go.mod
COPY        services/   services/

ENV         CGO_ENABLED=0
ENV         GOOS=linux
ENV         GOARCH=$arch

RUN         go mod tidy
RUN         go mod vendor
RUN         go build -o rest                    services/rest/rest/server.go
RUN         go build -o users-rest              services/rest/users/server.go
RUN         go build -o posts-rest              services/rest/posts/server.go
RUN         go build -o comments-rest           services/rest/comments/server.go
RUN         go build -o users-subgraph          services/graphql-subgraphs/users/server.go
RUN         go build -o posts-subgraph          services/graphql-subgraphs/posts/server.go
RUN         go build -o comments-subgraph       services/graphql-subgraphs/comments/server.go
RUN         go build -o notifications-subgraph  services/graphql-subgraphs/notifications/server.go
