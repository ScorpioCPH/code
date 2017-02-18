FROM jpetazzo/dind:latest

gcc for cgo
RUN apk update && apk add g++ \
        gcc \
        libc6-dev \
        make \
        pkg-config

ENV GOLANG_VERSION 1.8
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 53ab94104ee3923e228a2cb2116e5e462ad3ebaeea06ff04463479d7f12d27ca

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
    && tar -C /usr/local -xzf golang.tar.gz \
    && rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

COPY go-wrapper /usr/local/bin/

#####
# FROM docker:dind

# MAINTAINER CPH <scorpiocph@gmail.com>

# RUN apk --no-cache add bash go bzr git mercurial subversion openssh-client ca-certificates gcc musl-dev

# RUN apk --no-cache add openldap-dev libldap libsasl

# RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
# ENV GOPATH /go
# ENV PATH /go/bin:$PATH
# WORKDIR /go
