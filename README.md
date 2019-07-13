## API server

https://stage3-4eayrshuaa-uc.a.run.app

https://stage3-4eayrshuaa-uc.a.run.app/swagger/index.html


### デプロイ

```terminal
gcloud config set project ecsite-242111
gcloud builds submit --tag gcr.io/ecsite-242111/github.com/techforward/ecscript_server/stage
gcloud beta run deploy --image gcr.io/ecsite-242111/github.com/techforward/ecscript_server/stage --platform managed
stage3
```
https://cloud.google.com/cloud-build/docs/securing-builds/use-encrypted-secrets-credentials?hl=ja

### Docs

Create Address `POST /addresses`
```
{
    ulid: "xxxxxxxxxxx",
    postcode: "0000000",
    address1: "abc, def, ghi",
    address2: "2-22-2"
}
```

Create Auth `GET /auth/{firebase_token}`
```
return JWT
```

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
