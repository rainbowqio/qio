FROM golang:alpine3.18
LABEL vendor="RainbowQ"
LABEL app="qio"
LABEL version="0.2.1"
WORKDIR /go/src/qio/
COPY . .
RUN go get .
RUN go build
