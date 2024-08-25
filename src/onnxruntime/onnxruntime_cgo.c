#include "onnxruntime_cgo.h"
#include <stdlib.h>
#include <stdio.h>

// Function to retrieve the OrtApi pointer
const OrtApi* getOrtApi_2() {
    return OrtGetApiBase()->GetApi(ORT_API_VERSION);
}

// Function to create the ONNX Runtime environment
OrtEnv* createEnv_2(const OrtApi* g_ort) {
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