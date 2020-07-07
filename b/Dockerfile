FROM golang:latest

LABEL maintainer="smith-30"

WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

WORKDIR /app/b

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8022

# Command to run the executable
CMD ["./main"]