FROM golang:latest

RUN apt-get update && apt-get install -y vim 
RUN apt-get update && apt-get install -y locales \ 
  && sed -i -e 's/# \(ja_JP.UTF-8\)/\1/' /etc/locale.gen \
  && locale-gen \
  && update-locale LANG=ja_JP.UTF-8

RUN mkdir /mountain-searcher

WORKDIR /mountain-searcher

ADD  . /mountain-searcher

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/aws/aws-sdk-go-v2/config
RUN go get github.com/aws/aws-sdk-go-v2/service/sesv2
RUN go get github.com/aws/aws-sdk-go-v2/service/sesv2/types
