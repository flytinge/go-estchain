FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.7 https://github.com/estchain/go-estchain && \
  (cd go-estchain && make geth) && \
  cp go-estchain/build/bin/geth /geth && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-estchain && rm -rf /var/cache/apk/*

EXPOSE 8665
EXPOSE 31666

ENTRYPOINT ["/geth"]
