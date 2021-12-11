FROM ubuntu:16.04

MAINTAINER Ivan Rybakov "rybakov39@gmail.com"

ENV GOPATH /srv
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH

ADD ./account /srv/bin/account

EXPOSE 8080

CMD account