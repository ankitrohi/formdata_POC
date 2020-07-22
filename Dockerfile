FROM golang:1.13.5

WORKDIR /go/src/formdata_POC/
ADD . .

RUN go build -o main .

ENTRYPOINT ["./main"]

