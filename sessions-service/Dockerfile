
FROM golang:1.23-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY cmd/ ./cmd/
COPY internal/ ./internal/


RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go


FROM alpine:latest


COPY --from=builder /main /main

EXPOSE 8080


CMD ["/main"]