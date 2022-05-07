FROM golang:1.17-alpine

RUN mkdir /template
WORKDIR /template

COPY go.mod go.sum ./
RUN go mod tidy
