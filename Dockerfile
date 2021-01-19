FROM golang:latest

WORKDIR /go/src/github.com/cloudydong/scrapertestgo

COPY . .

CMD ["go","run","main.go"]