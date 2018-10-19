FROM golang:1.10-alpine3.7 as builder

WORKDIR /go/src

RUN apk update && apk add git

ADD src .

RUN go get github.com/astaxie/beego \
    github.com/beego/bee \
    github.com/olivere/elastic

RUN go build -o /go/bin/example.go example

EXPOSE 8080

ENTRYPOINT bee run example

FROM alpine:latest AS production

WORKDIR /opt/app

COPY --from=builder /go/bin/example.go .

RUN chmod +x example.go

ENTRYPOINT ./example.go