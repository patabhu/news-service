FROM golang:1.24-alpine

WORKDIR /app

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN go build -o app ./cmd/api

EXPOSE 8080

CMD ["./app"]
