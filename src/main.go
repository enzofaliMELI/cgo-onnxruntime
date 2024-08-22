package main

/*
#cgo CFLAGS: -I/opt/homebrew/Cellar/onnxruntime/1.17.1/include/onnxruntime
#cgo LDFLAGS: -L/opt/homebrew/Cellar/onnxruntime/1.17.1/lib -lonnxruntime -L. -lmyfuncs -lrunonnx

#include <stdio.h>
#include <stdlib.h>
#include <onnxruntime_c_api.h>
#include "myfuncs.h"
#include "runonnx.h"

*/
import "C"
import "unsafe"

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
}
