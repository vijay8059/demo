FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./ # Uncomment if you have dependencies

COPY . .

RUN go build -o demo-app main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/demo-app .

EXPOSE 8080

CMD ["./demo-app"]
