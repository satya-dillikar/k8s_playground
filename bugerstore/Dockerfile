FROM golang:1.16-alpine as builder

RUN mkdir -p /burgerstore/

WORKDIR /burgerstore

COPY . .
ENV GO111MODULE=on
# RUN go install sigs.k8s.io/json/internal/golang/encoding/json
RUN go mod download
RUN CGO_ENABLED=0 go build -a -o bin/server_main ./*.go

FROM alpine:3.14

WORKDIR /home/app

COPY --from=builder /burgerstore/bin/server_main .

CMD ["./server_main"]
