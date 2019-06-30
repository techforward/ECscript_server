FROM golang:1.12 as build_env

# Set GOPATH/GOROOT environment variables
RUN mkdir -p /go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN apt-get update
RUN apt-get upgrade -y


# Set up app
WORKDIR /go/src/github.com/techforward/ECscript_server
ADD . .

RUN go get -d -v ./...
RUN go install -v ./...

# ビルド処理 ビルドするときに必要なので環境変数をセットする。
ENV CGO_ENABLED 0
# 実際にビルド
RUN go build -v -o api_server /go/src/github.com/techforward/ECscript_server/main.go
RUN ls /go/src/github.com/techforward/ECscript_server

# 実行環境 alpineは軽量イメージ
FROM alpine:latest
WORKDIR /root/
# build_envステージからバイナリファイルだけをコピーしてくる。
COPY --from=build_env /go/src/github.com/techforward/ECscript_server/api_server .
RUN ls

EXPOSE 1323

# 起動時に実行
CMD ["./api_server"]