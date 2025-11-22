FROM golang:1.21.6-bookworm

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 6000

CMD [ "api" ]
