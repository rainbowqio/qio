FROM golang:alpine3.13
LABEL vendor="RainbowQ"
LABEL app="qio"
LABEL version="0.1.2"
WORKDIR /go/src/qio/
COPY . .
RUN go get .
RUN go build
