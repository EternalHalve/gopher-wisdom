    FROM golang:1.26-alpine AS builder

    RUN apk add --no-cache gcc musl-dev

    WORKDIR /src

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .

    RUN CGO_ENABLED=1 GOOS=linux go build -o /bin/app ./cmd/gopher-wisdom/main.go

    FROM alpine:latest
    RUN apk add --no-cache ca-certificates libc6-compat

    WORKDIR /app

    COPY --from=builder /bin/app .

    RUN mkdir -p /app/data

    EXPOSE 8080

    CMD ["./app"]