FROM golang:1.22-alpine AS builder

COPY . .
RUN go mod download
RUN go build -v -o /usr/local/bin/remedio cmd/main.go

FROM alpine
WORKDIR /root
COPY --from=builder /usr/local/bin/remedio /usr/local/bin/remedio

EXPOSE 8080

CMD ["remedio"]