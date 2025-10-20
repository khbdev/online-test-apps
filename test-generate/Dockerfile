FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o test-generation-servis ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/test-generation-servis .
COPY .env . 

EXPOSE 50054

CMD ["./test-generation-servis"]