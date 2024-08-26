#include "onnxruntime_cgo.h"
#include <stdlib.h>
#include <stdio.h>

// Function to retrieve the OrtApi pointer
const OrtApi* getOrtApi(char** error_message) {
    const OrtApi* api = OrtGetApiBase()->GetApi(ORT_API_VERSION);
    if (api == NULL) {
        const char* error_msg = "Failed to retrieve OrtApi pointer.";
        *error_message = strdup(error_msg);
    }
    return api;
}

// Function to create the ONNX Runtime environment
OrtEnv* createEnv(const OrtApi* api, char** error_message) {
    OrtEnv* env;
    OrtStatus* status= api->CreateEnv(ORT_LOGGING_LEVEL_WARNING, "test", &env);
    if (status != NULL) {
        printf("Error: %s\n", api->GetErrorMessage(status));
        const char* error_msg = api->GetErrorMessage(status);
        *error_message = strdup(error_msg);
        api->ReleaseStatus(status);
        return NULL;
    }

    return env;
}

// Function to release the ONNX Runtime environment
void releaseEnv(const OrtApi* api, OrtEnv* env) {
    if (env != NULL) {
        api->ReleaseEnv(env);
    }
}

// Function to create the ONNX Runtime session options
OrtSessionOptions* createSessionOptions(const OrtApi* api, char** error_message) {
    OrtSessionOptions* session_options;
    OrtStatus* status = api->CreateSessionOptions(&session_options);
    if (status != NULL) {
        const char* error_msg = api->GetErrorMessage(status);
        *error_message = strdup(error_msg);
        api->ReleaseStatus(status);
        return NULL;
    }

    return session_options;
}

// Function to release session options
void releaseSessionOptions(const OrtApi* api, OrtSessionOptions* session_options) {
    if (session_options != NULL) {
        api->ReleaseSessionOptions(session_options);
    }
}

// Function to create an ONNX Runtime session
OrtSession* createSession(const OrtApi* api, OrtEnv* env, const char* model_path, OrtSessionOptions* session_options, char** error_message) {
    OrtSession* session;
    OrtStatus* status = api->CreateSession(env, model_path, session_options, &session);
    if (status != NULL) {
        const char* error_msg = api->GetErrorMessage(status);
        *error_message = strdup(error_msg);
        api->ReleaseStatus(status);
        return NULL;
    }

    return session;
}

// Function to release an ONNX Runtime session
void releaseSession(const OrtApi* api, OrtSession* session) {
    if (session != NULL) {
        api->ReleaseSession(session);
    }
}

// Function to create an input tensor
OrtValue* createOrtTensor(const OrtApi* api, const float* input_data, size_t input_data_size, const int64_t* input_shape, size_t input_dim, char** error_message) {
    OrtMemoryInfo* memory_info;
    OrtStatus* status = api->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &memory_info);
    if (status != NULL) {
        const char* error_msg = api->GetErrorMessage(status);
        *error_message = strdup(error_msg);
        api->ReleaseStatus(status);
        return NULL;
    }

    OrtValue* tensor = NULL;
    status = api->CreateTensorWithDataAsOrtValue(memory_info, (void*)input_data, input_data_size, input_shape, input_dim, ONNX_TENSOR_ELEMENT_DATA_TYPE_FLOAT, &tensor);
    api->ReleaseMemoryInfo(memory_info);

    if (status != NULL) {
        const char* error_msg = api->GetErrorMessage(status);
        *error_message = strdup(error_msg);
        api->ReleaseStatus(status);
        return NULL;
    }

    return tensor;
}

// Function to release an ONNX Runtime tensor
void releaseOrtTensor(const OrtApi* api, OrtValue* tensor) {
    if (tensor != NULL) {
        api->ReleaseValue(tensor);
    }
}

// Function to run the model inference, including creating and returning the output tensor
OrtValue* runInference(const OrtApi* api, OrtSession* session, const char** input_names, const OrtValue* const* input_tensors, size_t input_count, const char** output_names, size_t output_count, char** error_message) {
    OrtValue* output_tensor = NULL;
    OrtStatus* status = api->Run(session, NULL, input_names, input_tensors, input_count, output_names, output_count, &output_tensor);

    if (status != NULL) {
        const char* error_msg = api->GetErrorMessage(status);
        *error_message = strdup(error_msg);
        api->ReleaseStatus(status);
        return NULL;
    }

    return output_tensor;
}

// Function to retrieve data from an output tensor
float* getTensorData(const OrtApi* api, OrtValue* tensor, char** error_message) {
    float* output_data;
    OrtStatus* status = api->GetTensorMutableData(tensor, (void**)&output_data);

    if (status != NULL) {
        const char* error_msg = api->GetErrorMessage(status);
        *error_message = strdup(error_msg);
        api->ReleaseStatus(status);
        return NULL;
    }

    return output_data;
}