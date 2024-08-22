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

	// Call the C function `runONNXRuntime` and capture the result
	tensorResult := C.runONNXRuntime(modelPath)
	defer C.free(unsafe.Pointer(tensorResult))

	// Convert the C pointer to a Go slice
	output := (*[10]C.float)(unsafe.Pointer(tensorResult))[:10:10]

	// Print the result in Go
	for i := 0; i < 10; i++ {
		fmt.Printf("output[%d] = %f\n", i, float32(output[i]))
	}

	fmt.Println("Go application finished.")
}
