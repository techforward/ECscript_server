## API server

### Docs
`go run main.go`
http://localhost:1323/swagger/index.html


### Build

docker build -t ecsite .
docker run -it --rm -p 1323:1323 ecsite

### Production

```terminal
server/api$ go build main.go 
server/api$ ./main &
```

### Development

```terminal
server/api$ go run main.go
```

### Testing

```terminal
server/api$ go test -v ./handler\
```
