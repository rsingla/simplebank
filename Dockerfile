# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o main .

EXPOSE 8080

CMD [ "/app/main" ]