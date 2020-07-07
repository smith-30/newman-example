FROM golang:latest

LABEL maintainer="smith-30"

WORKDIR /go/src/github.com/smith-30/cidind

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./a

# Expose port 8080 to the outside world
EXPOSE 8011

# Command to run the executable
CMD ["./main"]