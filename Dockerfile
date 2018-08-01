FROM golang:1.6

RUN wget http://storage.googleapis.com/kubernetes-release/release/v1.7.2/bin/linux/amd64/kubectl -O /usr/bin/kubectl && \
    chmod +x /usr/bin/kubectl

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app/

RUN go-wrapper download
RUN go-wrapper install

CMD ["app"]
