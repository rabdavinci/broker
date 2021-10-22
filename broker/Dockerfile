FROM golang:1.16-alpine AS build
RUN apk add --no-cache  git build-base
WORKDIR /build
COPY . .
# RUN go env -w GOPRIVATE=""
RUN go mod tidy
RUN go mod vendor
RUN go mod verify


