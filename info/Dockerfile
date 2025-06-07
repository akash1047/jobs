# Build stage
FROM golang:1.24.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o info main.go

FROM alpine:latest
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/info .
RUN chown appuser:appgroup info
USER appuser

ENTRYPOINT [ "./info" ]