FROM golang:1.14.4-alpine

LABEL maintainer="smith-30"

RUN apk update
RUN apk add --no-cache curl wget openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /go/src/github.com/smith-30/newman-example

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./a

# Expose port 8080 to the outside world
EXPOSE 8011

# Command to run the executable
CMD ["./main"]