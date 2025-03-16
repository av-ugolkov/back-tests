FROM golang:1.24.0-alpine3.21 as builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -o main

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /build/main /app/main

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]