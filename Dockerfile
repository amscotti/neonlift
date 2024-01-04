# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o neonlift .

# Final stage
FROM scratch
WORKDIR /root/
COPY --from=builder /app/neonlift .
ENTRYPOINT ["./neonlift"]