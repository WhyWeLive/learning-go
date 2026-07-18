FROM golang:alpine

WORKDIR /app/

COPY . /app/

RUN go mod download
RUN go build -o /app/server ./cmd/learning-go

CMD ["/app/server"]