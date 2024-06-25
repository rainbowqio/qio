FROM golang:alpine3.20
LABEL vendor="RainbowQ"
LABEL app="qio"
LABEL version="0.3.1"
WORKDIR /go/src/qio/
COPY . .
RUN go get .
RUN go build
