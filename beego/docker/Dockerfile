FROM golang:latest as builder

ADD src /go/src

RUN go get github.com/astaxie/beego \
    github.com/beego/bee \
    github.com/olivere/elastic

EXPOSE 9999

ENTRYPOINT bee run example
