# Stage 1: Build all Go binaries
FROM golang:1.23-alpine AS builder

LABEL maintainer="Zaid Albirawi"

RUN apk add --no-progress --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /go/src/github.com/zalbiraw/go-api-test-service

# Copy the go.mod and go.sum files for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY helpers/ helpers/
COPY services/ services/

# Build all the Go binaries
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/rest/rest/server                       services/rest/rest/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/rest/users/server                      services/rest/users/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/rest/posts/server                      services/rest/posts/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/rest/comments/server                   services/rest/comments/server.go

RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql/users/server                   services/graphql/users/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql/posts/server                   services/graphql/posts/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql/comments/server                services/graphql/comments/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql/notifications/server           services/graphql/notifications/server.go

RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql-subgraphs/users/server         services/graphql-subgraphs/users/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql-subgraphs/posts/server         services/graphql-subgraphs/posts/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql-subgraphs/comments/server      services/graphql-subgraphs/comments/server.go
RUN CGO_ENABLED=0 GO111MODULE=on go build -o services/graphql-subgraphs/notifications/server services/graphql-subgraphs/notifications/server.go

RUN CGO_ENABLED=1 GO111MODULE=on go build -tags musl -ldflags '-extldflags "-static"' -o services/kafka/notifications/producer services/kafka/notifications/producer.go

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Set the working directory
WORKDIR /go/src/github.com/zalbiraw/go-api-test-service

# Copy all the built binaries from the builder stage
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/helpers                                         /go/src/github.com/zalbiraw/go-api-test-service/helpers
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/rest/rest/server                       /go/src/github.com/zalbiraw/go-api-test-service/services/rest/rest/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/rest/users/server                      /go/src/github.com/zalbiraw/go-api-test-service/services/rest/users/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/rest/posts/server                      /go/src/github.com/zalbiraw/go-api-test-service/services/rest/posts/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/rest/comments/server                   /go/src/github.com/zalbiraw/go-api-test-service/services/rest/comments/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/graphql/users/server                   /go/src/github.com/zalbiraw/go-api-test-service/services/graphql/users/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/graphql/posts/server                   /go/src/github.com/zalbiraw/go-api-test-service/services/graphql/posts/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/graphql/comments/server                /go/src/github.com/zalbiraw/go-api-test-service/services/graphql/comments/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/graphql/notifications/server           /go/src/github.com/zalbiraw/go-api-test-service/services/graphql/notifications/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/server         /go/src/github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/server         /go/src/github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/comments/server      /go/src/github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/comments/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/server /go/src/github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/server
COPY --from=builder /go/src/github.com/zalbiraw/go-api-test-service/services/kafka/notifications/producer           /go/src/github.com/zalbiraw/go-api-test-service/services/kafka/notifications/producer
