FROM golang:1.23-alpine AS builder
LABEL maintainer="kietnguyen17052001@gmail.com"
ENV GOOS=linux \
    CGO_ENABLED=0 \
    GO111MODULE=on \
    GOFLAGS="-mod=vendor"

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor ./vendor
COPY . .

RUN go build -o main ./cmd/server/main.go

FROM alpine:latest AS production
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]
