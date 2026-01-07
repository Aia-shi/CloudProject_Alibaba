FROM golang:1.25.5-alpine3.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

FROM alpine:latest AS prod

WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE ${APP_PORT}

CMD ["./main"]
