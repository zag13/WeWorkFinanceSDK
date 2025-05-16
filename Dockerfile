# Set build arguments
ARG TARGETPLATFORM=linux/arm64
FROM --platform=${TARGETPLATFORM} golang:1.20.5

# Set environment variables
ENV CGO_ENABLED=1
ENV LD_LIBRARY_PATH=/app/C_sdk

# Set working directory
WORKDIR /app

# Copy go.mod first for better layer caching
COPY go.mod /app/
RUN go mod download

# Create C_sdk directory and copy SDK files based on architecture
RUN mkdir -p /app/C_sdk
COPY lib/arm64_C_sdk/* /app/C_sdk/

# Copy the rest of the source code
COPY finance_sdk.go /app/
COPY finance_sdk_test.go /app/

# Select correct SDK files based on architecture
RUN if [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
        rm -rf /app/C_sdk/* && \
        cp -r /app/lib/amd64_C_sdk/* /app/C_sdk/; \
    fi

# Build the project
RUN go build -o main .

# Run tests
# CMD ["tail", "-f", "/dev/null"]
CMD ["go", "test", "-v", "./..."]
