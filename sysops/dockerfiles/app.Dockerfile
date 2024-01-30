FROM golang:1.21-alpine as build-env

RUN apk add --update --upgrade build-base

WORKDIR /go/src

COPY . .
RUN GOOS=linux go build -o server

FROM alpine

WORKDIR /app

COPY --from=build-env /go/src/ .
EXPOSE 8000
ENTRYPOINT [ "./server" ]