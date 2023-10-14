FROM golang:1.20.7-bookworm

RUN apt-get update
RUN mkdir -p /home/echo-crud

WORKDIR /home/echo-crud

COPY . /home/echo-crud
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

ENV TZ=Asia/Tokyo