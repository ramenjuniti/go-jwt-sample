FROM golang:alpine

WORKDIR /go/src/server
COPY . .

RUN go build -o server

CMD ["./server"]