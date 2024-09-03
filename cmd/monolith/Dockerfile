FROM golang:1.21-alpine as build
RUN apk add --update --no-cache git

ENV GO111MODULE=on

WORKDIR /usr/src/app
RUN mkdir bin
COPY go.mod .
COPY go.sum .
COPY .env .
COPY configs/ configs/

RUN go mod download

COPY .. .
RUN go build -o bin/app cmd/monolith/main.go

FROM alpine:latest
RUN apk add ca-certificates
RUN apk add --update --no-cache curl
COPY --from=build /usr/src/app/bin/app /app
COPY --from=build /usr/src/app/.env /.env
COPY --from=build /usr/src/app/configs/ /configs/

CMD ["/app"]