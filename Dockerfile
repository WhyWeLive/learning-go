FROM golang:alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /server ./cmd/learning-go

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /server .

EXPOSE 8080

CMD ["./server"]