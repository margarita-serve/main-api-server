# syntax=docker/dockerfile:1

#FROM golang:1.18-alpine
FROM golang:1.16.3-stretch 

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

#RUN apk add build-base

RUN go build -o /koreserve

EXPOSE 8080

ENTRYPOINT [ "/koreserve" ]