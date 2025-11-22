FROM golang:1.22.0-bookworm

WORKDIR /app

RUN ls -al

COPY go.mod go.sum ./

RUN go mod download

COPY common common
COPY services services

RUN cd services/aggregator && \
    CGO_ENABLED=0 GOOS=linux go build -o /aggregator

CMD [ "/aggregator" ]
