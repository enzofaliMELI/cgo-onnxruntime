FROM golang:1.22.5
WORKDIR /app
COPY go.mod ./
COPY src/ ./src/
WORKDIR /app/src
ENV LD_LIBRARY_PATH=/app/src/onnxruntime/linux_aarch64/lib:$LD_LIBRARY_PATH
RUN go build -o main .
CMD ["./main"]