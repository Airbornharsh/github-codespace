# Stage 1: Build stage
FROM golang:1.21.3 as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go build -o service cmd/code-space/main.go

FROM alpine:3.12

WORKDIR /app
COPY --from=builder /app/service ./service

COPY ./data ./data

COPY ./pkg/dockerFiles ./pkg/dockerFiles

EXPOSE 5000

CMD ["./service"]
