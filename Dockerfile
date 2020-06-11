### -------- Stage 1: Build ---------- ###

FROM golang:1.12-alpine

RUN apk add --no-cache git

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

ENV PORT 8000

RUN go mod download

COPY ./src ./src

RUN go build -o ./out/go-api ./src
