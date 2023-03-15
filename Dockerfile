# Use an official Golang runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

RUN apk update && apk add --no-cache git postgresql-client

# Set environment variables
ENV POSTGRES_HOST postgres
ENV POSTGRES_PORT 5432
ENV POSTGRES_USER root
ENV POSTGRES_PASSWORD simplebankpass
ENV POSTGRES_DB postgres


# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]