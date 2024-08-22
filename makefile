# Compiler
CC = gcc

# Directories
SRC_DIR = src/onnxruntime
BUILD_DIR = build/onnxruntime

# Automatically find all .c and .h files
CFILES = $(wildcard $(SRC_DIR)/*.c)
HFILES = $(wildcard $(SRC_DIR)/*.h)

# Generate corresponding .o and .a file names
OBJECTS = $(patsubst $(SRC_DIR)/%.c, $(BUILD_DIR)/%.o, $(CFILES))
LIBS = $(patsubst $(SRC_DIR)/%.c, $(SRC_DIR)/lib%.a, $(CFILES))

# Default target: build libraries and run the Go program
all: $(LIBS)

# Compile C files into object files
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.c
	@mkdir -p $(BUILD_DIR)
	$(CC) $(CFLAGS) -c $< -o $@
	@echo "Compiled $<"

# Create static libraries from object files
$(SRC_DIR)/lib%.a: $(BUILD_DIR)/%.o
	ar rcs $@ $^

# Run the Go program
run: $(LIBS)
	cd src && go run main.go

# Clean up generated files
clean:
	rm -rf $(BUILD_DIR) $(LIBS)
	@echo "Cleaned up"