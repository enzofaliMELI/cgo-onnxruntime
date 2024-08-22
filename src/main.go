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

	// Call the C function `addNumbers`
	C.addNumbers(a, b)

	modelPath := C.CString("resources/naive_model.onnx")
	defer C.free(unsafe.Pointer(modelPath))

	C.runONNXRuntime(modelPath)

	fmt.Println("Go application finished.")
}
