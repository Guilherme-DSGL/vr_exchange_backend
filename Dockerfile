FROM golang:1.24.2-alpine3.21 AS builder

RUN apk update && apk upgrade

WORKDIR /app  

COPY . .

RUN go build -o engine ./app/

FROM alpine:latest

WORKDIR /app

EXPOSE 8080

COPY --from=builder /app/engine /app/
COPY app/.env . 

CMD /app/engine