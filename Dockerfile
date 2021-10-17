#build stage
FROM golang:alpine AS builder
COPY go.mod go.sum /go/src/github.com/AvestaBarzegar/statify-api/
WORKDIR /go/src/github.com/AvestaBarzegar/statify-api
RUN go mod download
COPY . /go/src/github.com/AvestaBarzegar/statify-api/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/statify-api github.com/AvestaBarzegar/statify-api/

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/AvestaBarzegar/statify-api//build/statify-api /usr/bin/statify-api
EXPOSE 8080
ENTRYPOINT ["/usr/bin/statify-api"]
