package main

/*
#cgo CFLAGS: -I./onnxruntime -I/usr/local/onnxruntime/include
#cgo LDFLAGS: -L./onnxruntime -lmyfuncs -lrunonnx -L/usr/local/onnxruntime/lib -lonnxruntime

#include "runonnx.h"
#include "myfuncs.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Define two integers to pass to the C function
	var a, b C.int
	a = 5
	b = 10

	// Call the C function `addNumbers` and capture the returned result
	result := C.addNumbers(a, b)

	// Print the result in Go
	fmt.Printf("The result of adding %d and %d is %d\n", int(a), int(b), int(result))

	modelPath := C.CString("resources/naive_model.onnx")
	defer C.free(unsafe.Pointer(modelPath))

	C.runONNXRuntime(modelPath)

	fmt.Println("Go application finished.")
}
