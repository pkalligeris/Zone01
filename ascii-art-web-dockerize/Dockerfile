# ============================================================
# Stage 1 — Build
# ============================================================
FROM golang:1.22-alpine AS builder

# Metadata (applied to every stage that inherits this label set)
LABEL maintainer="nbougial,pkallige,ipapigki"
LABEL version="1.0.0"
LABEL description="ASCII Art Web — a Go web server that renders ASCII art from text input"

# Set the working directory inside the container
WORKDIR /build

# Copy the Go module file first to leverage Docker layer caching.
# If go.mod hasn't changed, the cached layer is reused.
COPY go.mod ./

# Copy the rest of the source code
COPY . .

# Build a statically-linked binary (CGO disabled) so it runs
# on the minimal Alpine runtime without libc dependencies.
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/server ./cmd/ascii-art-web

# ============================================================
# Stage 2 — Runtime (minimal image)
# ============================================================
FROM alpine:3.19

LABEL maintainer="nbougial,pkallige,ipapigki"
LABEL version="1.0.0"
LABEL description="ASCII Art Web — a Go web server that renders ASCII art from text input"

# Install bash for auditor testing (docker exec -it <container> /bin/bash)
RUN apk add --no-cache bash

# Create a non-root user and group for security best practices
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy only what the runtime needs from the builder stage
COPY --from=builder /out/server ./server
COPY --from=builder /build/templates   ./templates
COPY --from=builder /build/assets      ./assets

# Change ownership so the non-root user can read everything
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose the port the server listens on
EXPOSE 8080

# Start the web server
CMD ["./server"]
