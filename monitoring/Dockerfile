FROM golang:1.23.0 AS builder

WORKDIR /app
COPY go.mod ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM gcr.io/distroless/base-debian11 AS final
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 3000

CMD ["/app/main"]