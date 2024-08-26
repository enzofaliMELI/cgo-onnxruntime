package main

import (
	"fmt"
	"github.com/enzofaliMELI/cgo-onnxruntime/src/onnxruntime"
	"reflect"
)

func main() {
	// Retrieve the OrtApi pointer
	api := onnxruntime.GetOrtApi()
	if api == nil {
		fmt.Println("Failed to get OrtApi")
		return
	}

	// Create the ONNX Runtime environment
	env := onnxruntime.CreateEnv(api)
	if env == nil {
		fmt.Println("Failed to create ONNX Runtime environment")
		return
	}
	fmt.Println(reflect.TypeOf(env))
	defer env.ReleaseEnv(api)

	// Create the Session Options
	options := onnxruntime.CreateSessionOptions(api)
	if options == nil {
		return
	}
	defer options.ReleaseSessionOptions(api)

	// Create the Session
	session := onnxruntime.CreateSession(api, env, "resources/naive_model.onnx", options)
	if session == nil {
		return
	}
	defer session.ReleaseSession(api)

	// Create the Input Tensor
	inputData := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	inputShape := []int64{1, 10} // Example shape
	tensor := onnxruntime.CreateTensor(api, inputData, inputShape)
	if tensor == nil {
		return
	}
	defer tensor.ReleaseTensor(api)

	fmt.Println("ONNX Runtime environment created successfully")

	fmt.Println("Go application finished.")
}
