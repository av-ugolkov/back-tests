FROM golang:1.24 as build

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o example

FROM scratch as image

COPY --from=build /app/example .

EXPOSE 2222

CMD ["/example"]
