FROM golang:1.25-alpine AS builder

RUN apk add --no-cache build-base ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM alpine:3.22

RUN apk --no-cache add ca-certificates sqlite tzdata

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/.env .

CMD ["./main"]