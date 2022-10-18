FROM golang:latest

WORKDIR /app
COPY . .
RUN go build -o gotchaPage *.go
CMD ["./gotchaPage"]
