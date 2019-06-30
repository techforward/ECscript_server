FROM golang:latest

WORKDIR /go/src/github.com/techforward/ECscript_server

COPY . .

RUN go get -d -v ...
RUN go install -v ...

RUN go build

CMD ["./api"]