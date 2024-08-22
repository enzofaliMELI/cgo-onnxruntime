
bash

```gcc -arch arm64 -I/opt/homebrew/Cellar/onnxruntime/1.17.1/include/onnxruntime -c runonnx.c -o runonnx.o```

```ar rcs librunonnx.a runonnx.o```

bash

```gcc -arch arm64 -c myfuncs.c -o myfuncs.o```

```ar rcs libmyfuncs.a myfuncs.o```


bash

```go run main.go```