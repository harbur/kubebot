FROM golang:1.9.0

RUN wget http://storage.googleapis.com/kubernetes-release/release/v1.9.11/bin/linux/amd64/kubectl -O /usr/bin/kubectl && \
    chmod +x /usr/bin/kubectl

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app/

RUN go get -d -v .
RUN go install -v .

CMD ["app"]
