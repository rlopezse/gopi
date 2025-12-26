# ---------- build stage ----------
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Copiamos solo los archivos de módulos (cache)
COPY go.mod ./
RUN go mod download

# Copiamos el resto (main.go, etc.)
COPY . .

# Compilamos el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -o api

# ---------- runtime stage ----------
FROM scratch

WORKDIR /

COPY --from=builder /app/api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
