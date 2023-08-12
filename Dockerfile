FROM golang:1.20.7-bookworm

RUN apt-get update
RUN mkdir -p /home/echo-crud

WORKDIR /home/echo-crud

COPY . /home/echo-crud
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/rubenv/sql-migrate/...@v1.5.2

ENV TZ=Asia/Tokyo