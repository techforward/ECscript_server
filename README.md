## API server

https://stage3-4eayrshuaa-uc.a.run.app

https://stage3-4eayrshuaa-uc.a.run.app/swagger/index.html

### Pre

https://cloud.google.com/cloud-build/docs/securing-builds/use-encrypted-secrets-credentials?hl=ja


### Docs

```
swag i
go run main.go
```
http://localhost:1323/swagger/index.html


### Build

docker build -t ecsite .
docker run -it --rm -p :1323:1323 ecsite

### Production

```terminal
server/api$ export MODE="prod"
server/api$ go build -v -o ECscript_server
server/api$ ./ECscript_server
```

### Development

```terminal
server/api$ go run main.go -mode local
```

### Testing

```terminal
server/api$ go test -v ./handler
```
