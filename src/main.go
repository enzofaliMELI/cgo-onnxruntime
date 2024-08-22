package main

/*
#cgo CFLAGS: -I/opt/homebrew/Cellar/onnxruntime/1.17.1/include/onnxruntime
#cgo LDFLAGS: -L/opt/homebrew/Cellar/onnxruntime/1.17.1/lib -lonnxruntime -L. -lmyfuncs

#include <stdio.h>
#include <stdlib.h>
#include <onnxruntime_c_api.h>
#include "myfuncs.h"

void runONNXRuntime() {
    // Initialize ONNX Runtime environment
    OrtEnv* env;
    OrtStatus* status;
    const OrtApi* g_ort = OrtGetApiBase()->GetApi(ORT_API_VERSION);

    status = g_ort->CreateEnv(ORT_LOGGING_LEVEL_WARNING, "test", &env);
    if (status != NULL) {
        printf("Error: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        return;
    }

    // Create a session options object
    OrtSessionOptions* session_options;
    status = g_ort->CreateSessionOptions(&session_options);
    if (status != NULL) {
        printf("Failed to create session options.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        g_ort->ReleaseEnv(env);
        return;
    }

    // Load the ONNX model
    const char* model_path = "resources/naive_model.onnx";  // Update with your model path
    OrtSession* session;
    status = g_ort->CreateSession(env, model_path, session_options, &session);
    if (status != NULL) {
        printf("Failed to load the ONNX model.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        g_ort->ReleaseSessionOptions(session_options);
        g_ort->ReleaseEnv(env);
        return;
    }

    // Clean up session options
    g_ort->ReleaseSessionOptions(session_options);

    // Prepare input tensor with shape [10]
    const int64_t input_shape[1] = { 10 };  // 1D tensor with 10 elements
    size_t input_tensor_size = 10 * sizeof(float);
    float input_tensor_data[10] = { 1.0f, 2.0f, 3.0f, 4.0f, 5.0f, 6.0f, 7.0f, 8.0f, 9.0f, 10.0f };  // Example input

    // Allocate memory info for the input tensor
    OrtMemoryInfo* memory_info;
    status = g_ort->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &memory_info);
    if (status != NULL) {
        printf("Failed to create memory info.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        g_ort->ReleaseSession(session);
        g_ort->ReleaseEnv(env);
        return;
    }

    // Create the input tensor
    OrtValue* input_tensor;
    status = g_ort->CreateTensorWithDataAsOrtValue(memory_info, input_tensor_data, input_tensor_size, input_shape, 1, ONNX_TENSOR_ELEMENT_DATA_TYPE_FLOAT, &input_tensor);
    if (status != NULL) {
        printf("Failed to create input tensor.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        g_ort->ReleaseMemoryInfo(memory_info);
        g_ort->ReleaseSession(session);
        g_ort->ReleaseEnv(env);
        return;
    }

    // Specify the input and output names
    const char* input_names[] = { "input" };
    const char* output_names[] = { "output" };

    // Prepare the output tensor (we'll allocate it later)
    OrtValue* output_tensor = NULL;

    // Run the model inference
    status = g_ort->Run(session, NULL, input_names, (const OrtValue*[]){ input_tensor }, 1, output_names, 1, &output_tensor);
    if (status != NULL) {
        printf("Failed to run model inference.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        g_ort->ReleaseValue(input_tensor);
        g_ort->ReleaseMemoryInfo(memory_info);
        g_ort->ReleaseSession(session);
        g_ort->ReleaseEnv(env);
        return;
    }

    // Retrieve and print the output tensor data
    float* output_data;
    status = g_ort->GetTensorMutableData(output_tensor, (void**)&output_data);
    if (status != NULL) {
        printf("Failed to get output tensor data.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
    } else {
        // Assume the output tensor is a 1D array of floats
        for (int i = 0; i < 10; i++) {
            printf("output[%d] = %f\n", i, output_data[i]);
        }
    }

    // Clean up
    g_ort->ReleaseValue(output_tensor);
    g_ort->ReleaseValue(input_tensor);
    g_ort->ReleaseMemoryInfo(memory_info);
    g_ort->ReleaseSession(session);
    g_ort->ReleaseEnv(env);

}

*/
import "C"

func main() {
	C.runONNXRuntime()

	// Define two integers to pass to the C function
	var a, b C.int
	a = 6
	b = 10

	// Call the C function `addNumbers`
	C.addNumbers(a, b)
}
