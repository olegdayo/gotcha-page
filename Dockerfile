FROM golang:latest

WORKDIR /app
COPY . .
RUN go build -o gotchaPage src/*.go
CMD ["./gotchaPage"]
