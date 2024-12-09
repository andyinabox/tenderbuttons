# Build Stage

# Write the dockerfile instructions
FROM golang:alpine AS BuildStage

# Set the working directory
WORKDIR /app

# Copy the Go source code
COPY . .

RUN go mod download

EXPOSE 8080

# Build the Go binary
RUN go build -o /main .

# Deploy Stage

FROM scratch

WORKDIR /

COPY --from=BuildStage /main /main

# Expose the port
EXPOSE 8080

# Run the Go binary
ENTRYPOINT ["/main"]