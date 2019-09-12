FROM golang:latest

WORKDIR /go
ADD dist/go /go
COPY data /go/data

CMD ["go", "run", "main.go"]

