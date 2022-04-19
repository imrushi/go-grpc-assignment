# Build Stage - Build go binary
FROM golang:1.17 as build

WORKDIR /gRPC
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gRPC_server main.go

# Package Stage - Package pdf_operations.bin
FROM alpine:3.11.3 

WORKDIR /gRPC
COPY --from=0 /gRPC/gRPC_server ./
ENTRYPOINT ["./gRPC_server"]
