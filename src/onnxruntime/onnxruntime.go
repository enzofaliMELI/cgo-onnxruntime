package onnxruntime

/*
#cgo CFLAGS: -I${SRCDIR}/onnxruntime -I/home/runner/onnxruntime/include
#cgo LDFLAGS: -L${SRCDIR}/onnxruntime -lrunonnx -L/home/runner/onnxruntime/lib -lonnxruntime -lonnxruntime_cgo

#include "onnxruntime_cgo.h"

*/
import "C"
import (
	"fmt"
)

// OnnxEnv wraps the C struct OrtEnv
type OnnxEnv struct {
	env *C.OrtEnv
}

// GetOrtApi retrieves the OrtApi pointer
func GetOrtApi() *C.OrtApi {
	api := C.getOrtApi()
	if api == nil {
		fmt.Println("Failed to get OrtApi")
		return nil
	}
	return api
}

// CreateEnv initializes the ONNX Runtime environment
func CreateEnv(api *C.OrtApi) *OnnxEnv {
	env := C.createEnv(api)
	if env == nil {
		fmt.Println("Failed to create ONNX Runtime environment")
		return nil
	}
	return &OnnxEnv{env: env}
}

// ExampleUsage Example usage in Go
func ExampleUsage() {
	api := GetOrtApi()
	if api == nil {
		return
	}

	env := CreateEnv(api)
	if env == nil {
		fmt.Println("Failed to create ONNX Runtime environment")
		return
	}

	fmt.Println("ONNX Runtime environment created successfully")
	// Remember to release the environment later when you're done with it.
}
