# Use an official Go image as the base image
FROM golang:1.22-alpine

# Install necessary build tools
RUN apk add --no-cache gcc g++ make curl

# Set the version of ONNX Runtime you want to install
ENV ONNXRUNTIME_VERSION=1.17.1

# Download and install the ONNX Runtime C library for ARM64
RUN curl -L https://github.com/microsoft/onnxruntime/releases/download/v${ONNXRUNTIME_VERSION}/onnxruntime-linux-aarch64-${ONNXRUNTIME_VERSION}.tgz \
    | tar xz -C /usr/local && \
    ln -s /usr/local/onnxruntime-linux-aarch64-${ONNXRUNTIME_VERSION} /usr/local/onnxruntime

# Set environment variables for the ONNX Runtime library
ENV CFLAGS="-I/usr/local/onnxruntime/include"
ENV LDFLAGS="-L/usr/local/onnxruntime/lib -lonnxruntime"

# Copy the source code and Makefile into the container
WORKDIR /app
COPY . .

# Build the application using Makefile
RUN make

# Run the Go application
CMD ["make", "run"]