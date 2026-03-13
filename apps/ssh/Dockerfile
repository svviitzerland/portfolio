# ── Build Stage ──────────────────────────────────────────────────────
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY data/ data/

WORKDIR /app/apps/ssh

COPY apps/ssh/go.mod apps/ssh/go.sum ./
RUN go mod download

COPY apps/ssh/ .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ssh-portfolio .

# ── Runtime Stage ───────────────────────────────────────────────────
FROM alpine:3.21

RUN apk --no-cache add ca-certificates

RUN adduser -D -g '' appuser

WORKDIR /app
RUN mkdir -p /app/.ssh && chown appuser:appuser /app/.ssh

USER appuser

COPY --from=builder /app/ssh-portfolio .

EXPOSE 2222

ENTRYPOINT ["./ssh-portfolio"]
