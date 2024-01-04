FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN mkdir -p /app

WORKDIR /app

COPY go.mod go.sum ./
        
RUN GOPROXY="https://goproxy.io" go mod download

