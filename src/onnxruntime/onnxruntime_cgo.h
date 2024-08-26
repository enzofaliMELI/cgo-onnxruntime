#ifndef ONNXRUNTIME_CGO_H
#define ONNXRUNTIME_CGO_H

#include <onnxruntime_c_api.h>

// Retrieves the OrtApi pointer for ONNX Runtime API access.
const OrtApi* getOrtApi();

// Initializes the ONNX Runtime environment.
OrtEnv* createEnv(const OrtApi* api);

// Releases the ONNX Runtime environment.
void releaseEnv(const OrtApi* api, OrtEnv* env);

// Creates session options for configuring an ONNX Runtime session.
OrtSessionOptions* createSessionOptions(const OrtApi* api);

// Releases the session options.
void releaseSessionOptions(const OrtApi* api, OrtSessionOptions* session_options);

// Creates an ONNX Runtime session for running a model.
OrtSession* createSession(const OrtApi* api, OrtEnv* env, const char* model_path, OrtSessionOptions* session_options);

// Releases the ONNX Runtime session.
void releaseSession(const OrtApi* api, OrtSession* session);

// Creates an input tensor for ONNX Runtime inference.
OrtValue* createOrtTensor(const OrtApi* api, const float* input_data, size_t input_data_size, const int64_t* input_shape, size_t input_dim);

// Releases an ONNX Runtime tensor.
void releaseOrtTensor(const OrtApi* api, OrtValue* tensor);

// Runs inference using the ONNX Runtime session and returns the output tensor(s).
OrtValue* runInference(const OrtApi* api, OrtSession* session, const char** input_names, const OrtValue* const* input_tensors, size_t input_count, const char** output_names, size_t output_count);

// Retrieves raw float data from an output tensor.
float* getTensorData(const OrtApi* api, OrtValue* tensor);

#endif // ONNXRUNTIME_CGO_H