package main

/*
#cgo CFLAGS: -I${SRCDIR}/onnxruntime -I/home/runner/onnxruntime/include
#cgo LDFLAGS: -L${SRCDIR}/onnxruntime -lmyfuncs -lrunonnx -L/home/runner/onnxruntime/lib -lonnxruntime

#include "runonnx.h"
#include "myfuncs.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	// Get the OrtApi pointer
	gOrt := C.getOrtApi()

	// Initialize ONNX Runtime environment
	env := C.createEnv(gOrt)
	if env == nil {
		fmt.Println("Failed to create ONNX Runtime environment")
		return
	}

	// Create session options
	sessionOptions := C.createSessionOptions(gOrt)
	if sessionOptions == nil {
		fmt.Println("Failed to create ONNX Runtime session options")
		return
	}

	modelPath := C.CString("resources/naive_model.onnx")
	defer C.free(unsafe.Pointer(modelPath))

	// Create session
	session := C.createSession(gOrt, env, modelPath, sessionOptions)
	if session == nil {
		fmt.Println("Failed to create ONNX Runtime session")
		return
	}

	// Prepare input tensor with shape [10]
	inputShape := [1]C.int64_t{10} // 1D tensor with 10 elements
	inputTensorSize := C.size_t(10 * C.sizeof_float)
	inputTensorData := [10]C.float{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0} // Example input

	// Create the input tensor
	inputTensor := C.createOrtTensor(gOrt, (*C.float)(unsafe.Pointer(&inputTensorData[0])), inputTensorSize, (*C.int64_t)(unsafe.Pointer(&inputShape[0])), 1)
	if inputTensor == nil {
		fmt.Println("Failed to create input tensor")
		return
	}

	// Specify the input and output names
	inputNames := []*C.char{C.CString("input")}
	defer C.free(unsafe.Pointer(inputNames[0]))
	outputNames := []*C.char{C.CString("output")}
	defer C.free(unsafe.Pointer(outputNames[0]))

	// Convert Go slices to C arrays
	inputNamesC := (**C.char)(unsafe.Pointer(&inputNames[0]))
	outputNamesC := (**C.char)(unsafe.Pointer(&outputNames[0]))

	// Run the model inference and get the output tensor
	outputTensor := C.runInference(gOrt, session, inputNamesC, &inputTensor, 1, outputNamesC, 1)
	if outputTensor == nil {
		fmt.Println("Failed to run model inference")
		return
	}

	// Retrieve the output tensor data
	outputData := C.getTensorData(gOrt, outputTensor)
	if outputData == nil {
		fmt.Println("Failed to get output tensor data")
		return
	}

	// Convert the C pointer to a Go slice
	output := (*[10]C.float)(unsafe.Pointer(outputData))[:10:10]

	// Print the output tensor data
	fmt.Println("Output Tensor Data:")
	for i := 0; i < 10; i++ {
		fmt.Printf("output[%d] = %f\n", i, float32(output[i]))
	}

	fmt.Println("Go application finished.")
}
