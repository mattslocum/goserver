FROM golang:1.15-alpine AS builder
LABEL maintainer="matt.slocum@gmail.com"

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/goserver


####### PROD CONTAINER #########
FROM scratch

WORKDIR /go/src/app
COPY --from=builder /go/bin/goserver /go/bin/goserver
EXPOSE 8080

ENTRYPOINT ["/go/bin/goserver"]
