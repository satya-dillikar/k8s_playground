ARG ALPINE_VERSION=3.14
ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as builder

RUN mkdir -p /go_app/

WORKDIR /go_app

COPY . .

ENV GO111MODULE=on

RUN go mod download

RUN CGO_ENABLED=0 go build -o bin/burgerstore ./*.go

FROM alpine:${ALPINE_VERSION}

WORKDIR /home/app

COPY --from=builder /go_app/bin/burgerstore .

ENTRYPOINT ["./burgerstore"]
