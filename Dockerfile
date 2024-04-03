# Use a smaller base image
FROM golang:alpine AS builder

# Set the working directory
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /snapp

# Final image
FROM alpine

# Copy the binary from the builder stage
COPY --from=builder /snapp /snapp

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["/snapp"]
