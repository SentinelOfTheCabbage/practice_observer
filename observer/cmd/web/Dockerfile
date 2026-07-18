FROM golang:1.26 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o executable ./cmd/web

FROM alpine:3.20
LABEL maintainer=SentinelOfTheCabbage
WORKDIR /app

COPY .env /app/.env

COPY --from=build-stage /app/executable ./executable
EXPOSE ${APP_PORT}
CMD ["./executable"]
