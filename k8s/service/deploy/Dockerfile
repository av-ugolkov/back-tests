FROM golang:1.24.0-alpine3.21 as builder

WORKDIR /build

COPY . .

RUN go build -o main

FROM scratch

COPY --from=builder /build/main /main

EXPOSE 5858

ENTRYPOINT [ "/main" ]

