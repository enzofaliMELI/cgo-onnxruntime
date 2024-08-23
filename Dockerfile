# Use an official Go image with Debian as the base image
FROM golang:1.22

# Install necessary build tools
RUN apt-get update && apt-get install -y gcc g++ make curl

# Set the version of ONNX Runtime you want to install
ENV ONNXRUNTIME_VERSION=1.17.1

# Create the directory structure
RUN mkdir -p /home/runner/onnxruntime

# Download and install the ONNX Runtime C library for x86_64 (64-bit Linux)
RUN curl -L https://github.com/microsoft/onnxruntime/releases/download/v${ONNXRUNTIME_VERSION}/onnxruntime-linux-aarch64-${ONNXRUNTIME_VERSION}.tgz \
    | tar xz -C /home/runner/onnxruntime --strip-components=1

# Set environment variables for the ONNX Runtime library to match your Go build expectations
ENV CFLAGS="-I/home/runner/onnxruntime/include"
ENV LDFLAGS="-L/home/runner/onnxruntime/lib -lonnxruntime"
ENV LD_LIBRARY_PATH=/home/runner/onnxruntime/lib

# Copy the source code and Makefile into the container
WORKDIR /app
COPY . .

# Build the application using Makefile
RUN make

# Run the Go application
CMD ["make", "run"]