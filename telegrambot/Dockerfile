FROM golang:1.13-alpine

WORKDIR /go/src/github.com/ravil23/lingualynda/telegrambot
COPY . .

RUN go get -v -d ./... \
    && go install -v ./...

ENTRYPOINT /go/bin/telegrambot
