# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy everything in the source directory to the working directory inside the container
COPY . .

# Ensure that the .env file is present and copied
COPY .env .env

# Run go mod tidy to clean up the dependencies
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/services cmd/services/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/migrate migrate/main.go


# Use a smaller image for the runtime environment
FROM alpine:3.16.0
WORKDIR /app/

RUN apk add --no-cache openssl
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Copy the binary and .env file from the builder image
COPY --from=builder /app/bin/services .
COPY --from=builder /app/bin/migrate .
COPY --from=builder /app/.env .
COPY --from=builder /app/startup.sh .  

RUN chmod +x ./startup.sh

# Expose the port the app runs on
EXPOSE 8000

# Run the application
# CMD ["./startup.sh"]
