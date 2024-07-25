FROM golang:1.22-alpine3.20 as builder

# Set the working directory to /app
WORKDIR /app

# Copy the Go source file into the container
COPY . .

# Build the Go program inside the container
RUN go build -o submodule-poc .

# Use a small Alpine Linux image as the base image for the final container
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the executable from the builder container
COPY --from=builder /app/submodule-poc .

# Create a user with a known UID/GID within range 10000-20000.
# This is required by Choreo to run the container as a non-root user.
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 10014 \
    "choreo"
# Use the above created unprivileged user
USER 10014

# Set the entrypoint to the executable
ENTRYPOINT ["./submodule-poc"]