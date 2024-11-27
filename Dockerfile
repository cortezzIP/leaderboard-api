FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o leaderboard-api cmd/main.go

FROM alpine:latest

COPY --from=builder /app/leaderboard-api .
COPY .env .

EXPOSE 8080

CMD ["./leaderboard-api"]