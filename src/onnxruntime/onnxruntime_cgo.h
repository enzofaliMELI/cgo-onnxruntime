#ifndef ONNXRUNTIME_CGO_H
#define ONNXRUNTIME_CGO_H

#include <onnxruntime_c_api.h>

// Function to retrieve the OrtApi pointer
const OrtApi* getOrtApi();

// Function to create the ONNX Runtime environment
OrtEnv* createEnv(const OrtApi* g_ort);


#endif