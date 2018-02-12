FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/estchain/go-estchain && \
  (cd go-estchain && make geth) && \
  cp go-estchain/build/bin/geth /geth && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-estchain

EXPOSE 8665
EXPOSE 31666

ENTRYPOINT ["/geth"]
