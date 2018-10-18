FROM golang:latest

ADD src /go/src

RUN go get github.com/astaxie/beego \
    github.com/beego/bee \
    go get github.com/olivere/elastic

EXPOSE 8080

ENTRYPOINT bee run tblox
