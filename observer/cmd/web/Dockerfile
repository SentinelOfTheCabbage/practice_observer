FROM golang:1.26

LABEL maintainer=SentinelOfTheCabbage

WORKDIR /app

COPY ../../go.mod .
RUN go mod download

COPY . .
RUN go build -o executable ./cmd/web

COPY .env /app
EXPOSE ${APP_PORT}
CMD ["./executable"]
