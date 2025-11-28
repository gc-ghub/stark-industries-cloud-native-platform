# -----------------------------
#  Stage 1: Build the Go binary
# -----------------------------
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o stark-industries main.go


# -----------------------------
#  Stage 2: Run with Distroless
# -----------------------------
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/stark-industries /app/stark-industries

USER nonroot:nonroot
EXPOSE 8080
CMD [ "./stark-industries" ]
