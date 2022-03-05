FROM golang:1.17-alpine

ENV TZ=Asia/Tokyo

ARG GNAME
ARG UNAME
ARG UID
ARG GID

WORKDIR /go/src/sake_io_auth

RUN apk add --update && \
    apk add git tzdata shadow && \
    addgroup -g ${GID} ${GNAME} && \
    adduser -D -H -u ${UID} -G ${GNAME} ${UNAME} && \
    go install github.com/cosmtrek/air@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install golang.org/x/lint/golint@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    chown -R go:go /go && \
    chown go:go /home

USER ${UNAME}:${GNAME}
EXPOSE 8090