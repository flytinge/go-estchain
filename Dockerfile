# Build Geth in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-estchain
RUN cd /go-estchain && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-estchain/build/bin/geth /usr/local/bin/

EXPOSE 8665 8666 31666 31666/udp 31667/udp
ENTRYPOINT ["geth"]
