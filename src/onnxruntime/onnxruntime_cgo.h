#ifndef ONNXRUNTIME_CGO_H
#define ONNXRUNTIME_CGO_H

#include <onnxruntime_c_api.h>

// Function to retrieve the OrtApi pointer
const OrtApi* getOrtApi();

// Function to create the ONNX Runtime environment
OrtEnv* createEnv(const OrtApi* g_ort);

OrtSessionOptions* createSessionOptions(const OrtApi* g_ort);

OrtSession* createSession(const OrtApi* g_ort, OrtEnv* env, const char* model_path, OrtSessionOptions* session_options);

OrtValue* createOrtTensor(const OrtApi* g_ort, const float* input_data, size_t input_data_size, const int64_t* input_shape, size_t input_dim);

OrtValue* runInference(const OrtApi* g_ort, OrtSession* session, const char** input_names, const OrtValue* const* input_tensors, size_t input_count, const char** output_names, size_t output_count);

float* getTensorData(const OrtApi* g_ort, OrtValue* tensor);

#endif