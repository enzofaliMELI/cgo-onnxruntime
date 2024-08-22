# Compiler
CC = gcc

# Directories
SRC_DIR = src/onnxruntime
BUILD_DIR = build/onnxruntime

# Files
CFILES = $(SRC_DIR)/runonnx.c $(SRC_DIR)/myfuncs.c
HFILES = $(SRC_DIR)/runonnx.h $(SRC_DIR)/myfuncs.h
OBJECTS = $(BUILD_DIR)/runonnx.o $(BUILD_DIR)/myfuncs.o
LIBS = $(SRC_DIR)/librunonnx.a $(SRC_DIR)/libmyfuncs.a

# Default target: build libraries and run the Go program
all: $(LIBS)

# Compile C files into object files
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.c $(SRC_DIR)/%.h
	@mkdir -p $(BUILD_DIR)
	$(CC) $(CFLAGS) -c $< -o $@
	@echo "Compiled $<"

# Create static libraries from object files
$(SRC_DIR)/librunonnx.a: $(BUILD_DIR)/runonnx.o
	ar rcs $@ $^

$(SRC_DIR)/libmyfuncs.a: $(BUILD_DIR)/myfuncs.o
	ar rcs $@ $^

# Run the Go program
run: $(LIBS)
	cd src && go run main.go

# Clean up generated files
clean:
	rm -rf $(BUILD_DIR) $(LIBS)
	@echo "Cleaned up"