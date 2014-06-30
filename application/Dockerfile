FROM ubuntu:14.04
MAINTAINER Brad Gignac <bgignac@bradgignac.com>

EXPOSE 8000
ENV GOPATH /usr/local/go
ENV GOBIN $GOPATH/bin
ENV GOSRC $GOPATH/src
ENV PATH $GOBIN:$PATH

RUN apt-get -y update
RUN apt-get -y install git golang

ADD main.go $GOSRC/hellogo/main.go
RUN go get hellogo

CMD hellogo
