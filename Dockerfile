FROM golang:1.27.1 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o /out/booking \
    ./cmd

FROM alpine:3
WORKDIR /app
COPY --from=builder /out/booking ./booking
EXPOSE 8080
ENTRYPOINT ["./booking"]