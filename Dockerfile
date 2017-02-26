FROM alpine:3.5

LABEL maintainer "fabiorphp@gmail.com"

ARG APP_NAME=myapp

RUN apk add --no-cache ca-certificates

ADD ./bin/$APP_NAME /app

CMD ["/app"]
