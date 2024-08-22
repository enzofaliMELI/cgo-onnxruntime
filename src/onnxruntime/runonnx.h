#ifndef RUNONNX_H
#define RUNONNX_H

#include <onnxruntime_c_api.h>

float* runONNXRuntime(const char* model_path);

#endif // RUNONNX_H