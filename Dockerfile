FROM golang:1.23.1 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /usr/src/hlb

COPY go.mod go.sum ./

RUN go mod download

# Copy build dependencies.
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg

# Build the api.
RUN go build -v -o ./bin/api ./cmd/api

FROM alpine:latest

WORKDIR /var/www/hlb

RUN apk add doas

# Copy the api.
COPY --from=builder /usr/src/hlb/bin/api .

# Copy configs.
COPY ./configs ./configs
