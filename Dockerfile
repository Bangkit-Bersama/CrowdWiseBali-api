FROM golang:1.23

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY cmd ./cmd
COPY internal ./internal
COPY service ./service

# Build
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/server

EXPOSE 8080

# Run
CMD ["/app/server"]
