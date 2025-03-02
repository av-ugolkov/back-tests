FROM golang:1.24.0-alpine3.21 as build

WORKDIR /build
COPY . .

RUN go mod init motel
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o main

FROM scratch as image

COPY --from=build /build/main .

EXPOSE 2222

CMD ["/main"]
