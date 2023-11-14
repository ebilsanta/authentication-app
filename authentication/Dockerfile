FROM golang:1.21-alpine3.17 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /service_executable

FROM alpine:3.17
COPY --from=builder /service_executable /service_executable
COPY *.env .

EXPOSE 8089
CMD ["/service_executable"]