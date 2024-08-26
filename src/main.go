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
	fmt.Println(reflect.TypeOf(options))
	defer options.ReleaseSessionOptions(api)

	fmt.Println("ONNX Runtime environment created successfully")

	fmt.Println("Go application finished.")
}
