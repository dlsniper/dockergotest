FROM golang:1.10.3-alpine3.7

ADD . /go/src/mypkg/

CMD ["go", "test", "-v", "-json", "mypkg/..."]