FROM golang:1.16-alpine

WORKDIR /go/src/

COPY ./main.go ./

RUN go mod init main
RUN go get github.com/skip2/go-qrcode
