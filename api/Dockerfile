FROM golang:1.16.8-alpine3.14

WORKDIR /api
COPY . .

RUN apk upgrade --update && \
    apk add --no-cache git

ENV TZ Asia/Tokyo

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["/api/main"]