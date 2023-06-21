# Build stage
FROM golang:1.20-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8088
RUN dos2unix /app/app.env
RUN chmod +x /app/start.sh
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]