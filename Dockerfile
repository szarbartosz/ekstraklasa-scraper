# build a binary
FROM golang:1.21.7-alpine AS builder

RUN apk update && apk add --no-cache 'git=~2'


ENV GO111MODULE=on
WORKDIR /scraper-app
COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /binary .

# buil a small image
FROM alpine:3

WORKDIR /
COPY --from=builder /binary /main

EXPOSE 8080

ENTRYPOINT ["/main"]