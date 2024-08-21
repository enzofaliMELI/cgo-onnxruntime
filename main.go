package main

/*
#cgo CFLAGS: -I/opt/homebrew/Cellar/onnxruntime/1.17.1/include/onnxruntime
#cgo LDFLAGS: -L/opt/homebrew/Cellar/onnxruntime/1.17.1/lib -lonnxruntime

#include <onnxruntime_c_api.h>
#include <stdlib.h>
#include <stdio.h>

const OrtApi* g_ort = NULL;

void runONNXRuntime() {
    OrtEnv* env;
    OrtStatus* status;

    g_ort = OrtGetApiBase()->GetApi(ORT_API_VERSION);
    if (g_ort == NULL) {
        printf("Failed to get ONNX Runtime API.\n");
        return;
    }

    // Initialize the ONNX Runtime environment
    status = g_ort->CreateEnv(ORT_LOGGING_LEVEL_WARNING, "test", &env);
    if (status != NULL) {
        const char* msg = g_ort->GetErrorMessage(status);
        printf("Failed to create ONNX Runtime environment: %s\n", msg);
        g_ort->ReleaseStatus(status);
        return;
    }
    printf("ONNX Runtime environment created successfully.\n");

    // Define model path (you should have a model.onnx file here)
    const char* model_path = "model.onnx";

    // Create session options
    OrtSessionOptions* session_options;
    g_ort->CreateSessionOptions(&session_options);

    // Load the model
    OrtSession* session;
    status = g_ort->CreateSession(env, model_path, session_options, &session);
    if (status != NULL) {
        const char* msg = g_ort->GetErrorMessage(status);
        printf("Failed to create ONNX Runtime session: %s\n", msg);
        g_ort->ReleaseStatus(status);
    } else {
        printf("ONNX model loaded successfully.\n");
        g_ort->ReleaseSession(session);
    }

    // Clean up
    g_ort->ReleaseSessionOptions(session_options);
    g_ort->ReleaseEnv(env);
}
*/
import "C"

func main() {
	C.runONNXRuntime()
}
