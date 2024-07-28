FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o pemtoder

FROM debian:bullseye-slim

WORKDIR /root/

COPY --from=builder /app/pemtoder .

CMD ["./pemtoder"]