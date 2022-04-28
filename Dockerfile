FROM golang:latest

RUN mkdir src/app
WORKDIR /src/app
COPY . .
RUN go build -o gotchaPage gotchaPage
CMD ["./gotchaPage"]
