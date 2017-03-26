FROM golang:1.8

RUN mkdir -p /go/src/github.com/limianwang/yo

WORKDIR /go/src/github.com/limianwang/yo

COPY . ./

RUN pwd

RUN go get -d -v ./...
RUN make build

EXPOSE 10001

CMD ["./yo"]
