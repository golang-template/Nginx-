# Use Golang image to build the service
FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go mod init service1 
RUN go build -o service1

# Create a minimal image for running the service
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/service1 .

CMD ["./service1"]
