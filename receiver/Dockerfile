FROM golang:1.19.3

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 7000

ENTRYPOINT [ "receiver" ]
