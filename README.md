## API server

https://test-4eayrshuaa-uc.a.run.app/swagger/index.html

### Docs

```
swag i
go run main.go
```
http://localhost:1323/swagger/index.html


### Build

docker build -t ecsite .
docker run -it --rm -p 127.0.0.1:1323:1323 ecsite

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
server/api$ go test -v ./handler\
```
