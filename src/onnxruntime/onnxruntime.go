package onnxruntime

/*
#cgo CFLAGS: -I${SRCDIR}/onnxruntime -I/home/runner/onnxruntime/include
#cgo LDFLAGS: -L${SRCDIR}/onnxruntime -L/home/runner/onnxruntime/lib -lonnxruntime

#include "onnxruntime_cgo.h"

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// OnnxEnv wraps the C struct OrtEnv
type OnnxEnv struct {
	env *C.OrtEnv
}

// OnnxSessionOptions wraps the C struct OrtSessionOptions
type OnnxSessionOptions struct {
	options *C.OrtSessionOptions
}

// OnnxSession wraps the C struct OrtSession
type OnnxSession struct {
	session *C.OrtSession
}

// OnnxTensor wraps the C struct OrtValue
type OnnxTensor struct {
	tensor *C.OrtValue
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

// CreateSession initializes the ONNX Runtime session
func CreateSession(api *C.OrtApi, env *OnnxEnv, modelPath string, options *OnnxSessionOptions) *OnnxSession {
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))

	session := C.createSession(api, env.env, cModelPath, options.options)
	if session == nil {
		fmt.Println("Failed to create ONNX Runtime session")
		return nil
	}
	return &OnnxSession{session: session}
}

// ReleaseSession releases the ONNX Runtime session
func (s *OnnxSession) ReleaseSession(api *C.OrtApi) {
	if s.session != nil {
		C.releaseSession(api, s.session)
		s.session = nil // Avoid double free
	}
}

// CreateTensor creates an ONNX Runtime tensor
func CreateTensor(api *C.OrtApi, inputData []float32, inputShape []int64) *OnnxTensor {
	inputDataSize := C.size_t(len(inputData) * int(unsafe.Sizeof(inputData[0])))
	inputDim := C.size_t(len(inputShape))

	tensor := C.createOrtTensor(api, (*C.float)(unsafe.Pointer(&inputData[0])), inputDataSize, (*C.int64_t)(unsafe.Pointer(&inputShape[0])), inputDim)
	if tensor == nil {
		fmt.Println("Failed to create ONNX Runtime tensor")
		return nil
	}
	return &OnnxTensor{tensor: tensor}
}

// ReleaseTensor releases an ONNX Runtime tensor
func (t *OnnxTensor) ReleaseTensor(api *C.OrtApi) {
	if t.tensor != nil {
		C.releaseOrtTensor(api, t.tensor)
		t.tensor = nil // Avoid double free
	}
}

// RunInference runs inference on the model and returns the output tensor
func RunInference(api *C.OrtApi, session *OnnxSession, inputNames []string, inputTensors []*OnnxTensor, outputNames []string) *OnnxTensor {
	// Convert input names to C strings
	cInputNames := make([]*C.char, len(inputNames))
	for i, name := range inputNames {
		cInputNames[i] = C.CString(name)
		defer C.free(unsafe.Pointer(cInputNames[i]))
	}

	// Convert output names to C strings
	cOutputNames := make([]*C.char, len(outputNames))
	for i, name := range outputNames {
		cOutputNames[i] = C.CString(name)
		defer C.free(unsafe.Pointer(cOutputNames[i]))
	}

	// Convert input tensors to C array
	cInputTensors := make([]*C.OrtValue, len(inputTensors))
	for i, tensor := range inputTensors {
		cInputTensors[i] = tensor.tensor
	}

	outputTensor := C.runInference(api, session.session, &cInputNames[0], &cInputTensors[0], C.size_t(len(inputTensors)), &cOutputNames[0], C.size_t(len(outputNames)))
	if outputTensor == nil {
		fmt.Println("Failed to run inference")
		return nil
	}
	return &OnnxTensor{tensor: outputTensor}
}

// GetTensorData retrieves data from an output tensor
func GetTensorData(api *C.OrtApi, tensor *OnnxTensor, size int) []float32 {
	data := C.getTensorData(api, tensor.tensor)
	if data == nil {
		fmt.Println("Failed to get tensor data")
		return nil
	}

	// Convert the C pointer to a Go slice
	output := (*[1 << 30]C.float)(unsafe.Pointer(data))[:size:size]

	// Convert the C float slice to a Go float32 slice
	goOutput := make([]float32, size)
	for i := 0; i < size; i++ {
		goOutput[i] = float32(output[i])
	}
	return goOutput
}
