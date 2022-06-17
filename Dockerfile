# syntax=docker/dockerfile:1
FROM golang:1.18-alpine
WORKDIR /app
COPY handlers /app/handlers
COPY go.mod ./
COPY go.sum ./
RUN apk update && apk add git
COPY *.go ./
RUN go build -o /server
EXPOSE 8080
CMD [ "/server" ]