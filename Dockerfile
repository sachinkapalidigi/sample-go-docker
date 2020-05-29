FROM golang:1.14-alpine

ENV GO111MODULE=on

WORKDIR /go/src/deploy-sample
COPY ./src/go.mod .
COPY ./src/go.sum .

RUN go mod download

COPY ./src .

RUN go build

EXPOSE 8081

CMD ["./src"]