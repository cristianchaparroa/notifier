# Step 1: Modules caching
FROM golang:1.21-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.21-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
# ARG GOARCH= // I'm on mac so i'm disabling this temporary
ARG CGO_ENABLED=0
ARG GOOS=linux
RUN apk add alpine-sdk && \
    go build -o /bin/notification ./cmd/notification/main.go
##
## Deploy
##
FROM gcr.io/distroless/base-debian10
COPY --from=builder /bin/notification /notification

ENV APP_PORT=8000
CMD ["/notification"]
