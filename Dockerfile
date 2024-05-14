FROM golang:1.22-alpine

WORKDIR /usr/src/app

COPY . .
RUN go mod download
RUN go build -v -o /usr/local/bin/remedio cmd/main.go

EXPOSE 8080

CMD ["remedio"]
