FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

FROM scratch

COPY --from=builder /app/cmd/ordersystem/.env /
COPY --from=builder /app/app /

CMD ["/app"]
