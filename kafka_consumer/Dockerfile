FROM golang:1.21.6-bookworm

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD [ "consumer" ]
