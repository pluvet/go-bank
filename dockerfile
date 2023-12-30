FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app
    
COPY go.mod go.sum .env ./
        
RUN go mod download

