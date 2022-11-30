FROM golang:1.18-alpine3.14 AS builder
WORKDIR /src/app


COPY . .
RUN go mod download
RUN go build -o /src/srv .

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /src/srv .

ENTRYPOINT [ "./srv" ]
