FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk update && apk add --no-cache git

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]