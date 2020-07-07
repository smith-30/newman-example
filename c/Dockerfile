FROM golang:latest

LABEL maintainer="smith-30"

WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

WORKDIR /app/c

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]