FROM golang:1.22.4-alpine AS builder

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o /app/main

EXPOSE 8080

CMD ["/app/main"]