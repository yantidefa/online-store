FROM golang:1.20-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod tidy
COPY ./ /usr/src/app/
RUN go build -ldflags -w -o ./bin/online-store .

FROM alpine:latest as run

WORKDIR /apps
COPY --from=builder /usr/src/app/bin /apps/
COPY .env /apps/

ENTRYPOINT ["/apps/online-store"]