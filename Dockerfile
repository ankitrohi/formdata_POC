FROM golang:1.13.5

WORKDIR /go/src
ADD . .

RUN go build -o main .

ENTRYPOINT ["./main"]

