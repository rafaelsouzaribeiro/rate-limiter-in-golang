FROM golang:1.26.3

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ./build/rate-limiter ./cmd/main.go

CMD ["./build/rate-limiter"]