FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
        
COPY go.mod go.sum main.go ./
        
RUN mkdir -p /app

WORKDIR /app
