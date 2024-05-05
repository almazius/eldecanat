FROM golang:1.22-alpine as build

WORKDIR /deck

COPY go.mod go.sum /deck/
COPY . /deck

RUN go build -v -o main eldeck/main/main.go