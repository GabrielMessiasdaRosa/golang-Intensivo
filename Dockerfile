# Build stage
FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./cmd/order/main.go

# Final stage
FROM scratch

COPY --from=builder /app/main /

CMD ["/main"]