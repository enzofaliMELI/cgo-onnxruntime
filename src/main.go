package main

import (
	"fmt"
	"github.com/enzofaliMELI/cgo-onnxruntime/src/onnxruntime"
	"log"
	"os"
)

func main() {
	// Assuming the ONNX library is located in the project directory under "src/onnxruntime/linux_aarch64/lib"
	libDir := "/Users/efaliveni/Desktop/projects/cgo-onnxruntime/src/onnxruntime/linux_aarch64/lib"

	// Set the LD_LIBRARY_PATH environment variable to include the directory containing the ONNX Runtime library
	err := os.Setenv("LD_LIBRARY_PATH", libDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to set LD_LIBRARY_PATH: %v", err))
	}

	// Retrieve the OrtApi pointer
	api, err := onnxruntime.GetOrtApi()
	if err != nil {
		log.Fatalf("Error retrieving OrtApi pointer: %v", err)
	}

	// Create the ONNX Runtime environment
	env, err := onnxruntime.CreateEnv(api)
	if err != nil {
		log.Fatalf("Error creating ONNX Runtime environment: %v", err)
	}
	defer env.ReleaseEnv(api)

	// Create the Session Options
	sessionOptions, err := onnxruntime.CreateSessionOptions(api)
	if err != nil {
		log.Fatalf("Error creating ONNX Runtime session options: %v", err)
	}
	defer sessionOptions.ReleaseSessionOptions(api)

	// Create the Session
	session, err := onnxruntime.CreateSession(api, env, "resources/naive_model.onnx", sessionOptions)
	if err != nil {
		log.Fatalf("Error creating ONNX Runtime session: %v", err)
	}
	defer session.ReleaseSession(api)

	// Create the Input Tensor
	inputShape := []int64{10}                                                 // 1D tensor with 10 elements
	inputData := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0} // Example input
	tensor, err := onnxruntime.CreateTensor(api, inputData, inputShape)
	if err != nil {
		log.Fatalf("Error creating ONNX Runtime tensor: %v", err)
	}
	defer tensor.ReleaseTensor(api)

	inputNames := []string{"input"}
	outputNames := []string{"output"}

	outputTensor, err := onnxruntime.RunInference(api, session, inputNames, []*onnxruntime.OnnxTensor{tensor}, outputNames)
	if err != nil {
		log.Fatalf("Error running ONNX Runtime inference: %v", err)
	}
	defer outputTensor.ReleaseTensor(api)

	// Retrieve the Output Data
	outputData, err := onnxruntime.GetTensorData(api, outputTensor, 10) // Specify the size of the output tensor
	if err != nil {
		log.Fatalf("Error retrieving tensor data: %v", err)
	}

	// Print the output data
	fmt.Println("Output Tensor Data:")
	for i, val := range outputData {
		fmt.Printf("output[%d] = %f\n", i, val)
	}

	fmt.Println("Go application finished.")
}
