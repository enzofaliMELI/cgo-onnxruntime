package onnxruntime

/*
#cgo CFLAGS: -I${SRCDIR}/onnxruntime -I/home/runner/onnxruntime/include
#cgo LDFLAGS: -L${SRCDIR}/onnxruntime -L/home/runner/onnxruntime/lib -lonnxruntime

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

// OnnxSessionOptions wraps the C struct OrtSessionOptions
type OnnxSessionOptions struct {
	options *C.OrtSessionOptions
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

// ReleaseEnv releases the ONNX Runtime environment
func (e *OnnxEnv) ReleaseEnv(api *C.OrtApi) {
	if e.env != nil {
		C.releaseEnv(api, e.env)
		e.env = nil // Avoid double free
	}
}

// CreateSessionOptions initializes the ONNX Runtime session options
func CreateSessionOptions(api *C.OrtApi) *OnnxSessionOptions {
	options := C.createSessionOptions(api)
	if options == nil {
		fmt.Println("Failed to create ONNX Runtime session options")
		return nil
	}
	return &OnnxSessionOptions{options: options}
}

// ReleaseSessionOptions releases the ONNX Runtime session options
func (o *OnnxSessionOptions) ReleaseSessionOptions(api *C.OrtApi) {
	if o.options != nil {
		C.releaseSessionOptions(api, o.options)
		o.options = nil // Avoid double free
	}
}
