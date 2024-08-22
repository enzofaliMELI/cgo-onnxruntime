
Local

```gcc -arch arm64 -I/opt/homebrew/Cellar/onnxruntime/1.17.1/include/onnxruntime -c runonnx.c -o runonnx.o```

```ar rcs librunonnx.a runonnx.o```

```gcc -arch arm64 -c myfuncs.c -o myfuncs.o```

```ar rcs libmyfuncs.a myfuncs.o```


```go run main.go```


Docker

```docker build -t my-cgo-app .```

```docker run --name my-cgo-app-container my-cgo-app```

```docker logs my-cgo-app-container```


Docker-compose

```docker-compose up```

```docker-compose up --build```

```docker-compose down```


onnxruntime C

https://onnxruntime.ai/docs/get-started/with-c.html