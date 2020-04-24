FROM golang:1.14-alpine3.11 as dev

RUN apk add alpine-sdk zsh

env GO=/usr/local/go/bin/go

env BUILD_DIR=/build/github.com/cuberat/go-log

RUN mkdir -p ${BUILD_DIR}

COPY ./ ${BUILD_DIR}/

RUN cd ${BUILD_DIR}/test_utils/test_syslog && \
    go build -a && \
    cp test_syslog /usr/local/bin/

# ------------------------------------------------------------------------------
FROM dev as prod

CMD ["/sbin/syslogd", "-n", "-O", "/dev/stderr"]
