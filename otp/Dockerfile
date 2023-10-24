FROM golang:1.21-alpine3.17
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /docker-gs-ping

EXPOSE 8089
CMD ["/docker-gs-ping"]