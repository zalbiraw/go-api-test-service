FROM        golang:1.17.6-alpine

MAINTAINER  Zaid Albirawi

WORKDIR     /go/src/github.com/zalbiraw/go-api-test-service

COPY        go.mod      go.mod
COPY        helpers/    helpers/
COPY        services/   services/

RUN         go mod tidy
RUN         go mod vendor
RUN         CGO_ENABLED=0 GO111MODULE=on go build -o services/rest/rest                        services/rest/rest/server.go
RUN         CGO_ENABLED=0 GO111MODULE=on go build -o services/rest/users                       services/rest/users/server.go
RUN         CGO_ENABLED=0 GO111MODULE=on go build -o services/rest/posts                       services/rest/posts/server.go
RUN         CGO_ENABLED=0 GO111MODULE=on go build -o services/rest/comments                    services/rest/comments/server.go
RUN         CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql-subgraphs/users          services/graphql-subgraphs/users/server.go
RUN         CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql-subgraphs/posts          services/graphql-subgraphs/posts/server.go
RUN         CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql-subgraphs/comments       services/graphql-subgraphs/comments/server.go
RUN         CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql-subgraphs/notifications  services/graphql-subgraphs/notifications/server.go
