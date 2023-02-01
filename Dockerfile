FROM alpine:latest

COPY goet /usr/local/bin/goet

ENTRYPOINT ["/usr/local/bin/goet"]
