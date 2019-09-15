############################
# STEP 1 build executable binary
############################
FROM golang:1.12-alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add --no-cache ca-certificates
WORKDIR /src
COPY go.mod go.mod
COPY go.sum go.sum
# Using go mod.
RUN go mod download
RUN go mod verify
COPY . .
# Build the binary.
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -ldflags="-w -s" -o /app/backend
############################
# STEP 2 build a small image
############################
FROM scratch
# Import the user and group files from the builder.
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy our static executable.
COPY --from=builder /app/backend /app
# Use an unprivileged user.
USER nobody:nobody
# Run the hello binary.
ENTRYPOINT ["/app"]
