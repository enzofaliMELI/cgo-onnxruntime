# Start with a base image that has Go installed
FROM golang:1.20

# Set environment variables for ONNX Runtime version
ENV ORT_VERSION=1.17.1

# Install necessary tools and dependencies
RUN apt-get update && apt-get install -y \
    curl \
    build-essential \
    cmake \
    libprotobuf-dev \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

# Download and install ONNX Runtime
RUN curl -L https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VERSION}/onnxruntime-linux-x64-${ORT_VERSION}.tgz -o onnxruntime-linux.tgz \
    && tar -xzvf onnxruntime-linux.tgz \
    && mv onnxruntime-linux-x64-${ORT_VERSION} /opt/onnxruntime \
    && rm onnxruntime-linux.tgz

# Set environment variables for ONNX Runtime
ENV CGO_CFLAGS="-I/opt/onnxruntime/include"
ENV CGO_LDFLAGS="-L/opt/onnxruntime/lib -lonnxruntime"
ENV LD_LIBRARY_PATH="/opt/onnxruntime/lib:$LD_LIBRARY_PATH"

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code
COPY . .

# Build the Go application
RUN go build -o myapp

# Command to run the Go application
CMD ["./myapp"]