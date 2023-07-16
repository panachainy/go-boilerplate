FROM golang:1.20-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd

FROM alpine:latest
RUN apk add --update --no-cache ca-certificates git

COPY --from=builder ["/build/api", "/"]

EXPOSE 8080

ENTRYPOINT ["/api"]
