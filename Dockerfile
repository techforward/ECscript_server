# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/techforward/ECscript_server
COPY . .

RUN go get -v ./...

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
ENV MODE production
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ECscript_server

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/techforward/ECscript_server/ECscript_server ECscript_server

ENV PORT 1323
EXPOSE 1323

# Run the web service on container startup.
CMD ["./ECscript_server"]

# gcloud builds submit --tag gcr.io/ecsite-242111/github.com/techforward/test

# gcloud beta run deploy --image gcr.io/ecsite-242111/github.com/techforward/test --platform managed