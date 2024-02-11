FROM golang:1.22.0-bookworm

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD [ "consumer" ]
