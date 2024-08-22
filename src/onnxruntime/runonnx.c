#include <stdio.h>
#include <stdlib.h>
#include "runonnx.h"

// Function to retrieve the OrtApi pointer
const OrtApi* getOrtApi() {
    return OrtGetApiBase()->GetApi(ORT_API_VERSION);
}

// Function to create the ONNX Runtime environment
OrtEnv* createEnv(const OrtApi* g_ort) {
    OrtEnv* env;
    OrtStatus* status;

    status = g_ort->CreateEnv(ORT_LOGGING_LEVEL_WARNING, "test", &env);
    if (status != NULL) {
        printf("Error: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        return NULL;
    }

    return env;
}

// Function to create the ONNX Runtime session options
OrtSessionOptions* createSessionOptions(const OrtApi* g_ort) {
    OrtSessionOptions* session_options;
    OrtStatus* status;

    status = g_ort->CreateSessionOptions(&session_options);
    if (status != NULL) {
        printf("Failed to create session options.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        return NULL;
    }

    return session_options;
}

// Function to create the ONNX Runtime session
OrtSession* createSession(const OrtApi* g_ort, OrtEnv* env, const char* model_path, OrtSessionOptions* session_options) {
    OrtSession* session;
    OrtStatus* status;

    status = g_ort->CreateSession(env, model_path, session_options, &session);
    if (status != NULL) {
        printf("Failed to load the ONNX model.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        return NULL;
    }

    return session;
}

OrtValue* createOrtTensor(const OrtApi* g_ort, const float* input_data, size_t input_data_size, const int64_t* input_shape, size_t input_dim) {
    OrtMemoryInfo* memory_info;
    OrtValue* input_tensor;
    OrtStatus* status;

    // Allocate memory info for the input tensor
    status = g_ort->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &memory_info);
    if (status != NULL) {
        printf("Failed to create memory info.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        return NULL;
    }

    // Create the input tensor
    status = g_ort->CreateTensorWithDataAsOrtValue(memory_info, (void*)input_data, input_data_size, input_shape, input_dim, ONNX_TENSOR_ELEMENT_DATA_TYPE_FLOAT, &input_tensor);
    if (status != NULL) {
        printf("Failed to create input tensor.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        g_ort->ReleaseMemoryInfo(memory_info);
        return NULL;
    }

    // Release memory info after tensor creation
    g_ort->ReleaseMemoryInfo(memory_info);

    return input_tensor;
}

// Function to run the model inference, including creating and returning the output tensor
OrtValue* runInference(const OrtApi* g_ort, OrtSession* session, const char** input_names, const OrtValue* const* input_tensors, size_t input_count, const char** output_names, size_t output_count) {
    OrtStatus* status;
    OrtValue* output_tensor = NULL;

    // Run the model inference
    status = g_ort->Run(session, NULL, input_names, input_tensors, input_count, output_names, output_count, &output_tensor);
    if (status != NULL) {
        printf("Failed to run model inference.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        return NULL;
    }

    return output_tensor;  // Return the output tensor
}

// Function to retrieve and return the tensor data
float* getTensorData(const OrtApi* g_ort, OrtValue* tensor) {
    float* output_data = NULL;
    OrtStatus* status;

    // Retrieve the output tensor data
    status = g_ort->GetTensorMutableData(tensor, (void**)&output_data);
    if (status != NULL) {
        printf("Failed to get output tensor data.\nError: %s\n", g_ort->GetErrorMessage(status));
        g_ort->ReleaseStatus(status);
        return NULL;
    }

    return output_data;  // Return the pointer to the data
}