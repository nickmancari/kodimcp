FROM golang:1.23-alpine AS build

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN go build -o /kodimcp ./cmd/kodimcp

FROM alpine:3.20
COPY --from=build /kodimcp /usr/local/bin/kodimcp
ENTRYPOINT ["kodimcp"]
