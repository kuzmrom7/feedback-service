FROM golang:1.15-alpine3.12 AS builder

RUN go version

COPY . /github.com/kuzmrom7/feedback-service/
WORKDIR /github.com/kuzmrom7/feedback-service/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/main ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/kuzmrom7/feedback-service/.bin/bot .
COPY --from=0 /github.com/kuzmrom7/feedback-service/configs configs/

EXPOSE 80

CMD ["./bot"]