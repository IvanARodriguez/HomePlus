# Base image
FROM golang:1.22-alpine AS base

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download

COPY . .

FROM base AS dev

EXPOSE 8000

CMD [ "air", "-c", ".air.toml" ]