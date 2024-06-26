FROM golang:1.19-alpine as stage1

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
