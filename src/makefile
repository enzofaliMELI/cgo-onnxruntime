# Compiler and flags
CC = gcc
CFLAGS = -arch arm64 -I/opt/homebrew/Cellar/onnxruntime/1.17.1/include/onnxruntime -Ionnx

# Target libraries and object files
LIBS = onnx/librunonnx.a onnx/libmyfuncs.a
OBJECTS = onnx/runonnx.o onnx/myfuncs.o

# Default target: build libraries and run the Go program
all: $(LIBS) run

# Rule to create librunonnx.a
onnx/librunonnx.a: onnx/runonnx.o
	ar rcs $@ $^

# Rule to create libmyfuncs.a
onnx/libmyfuncs.a: onnx/myfuncs.o
	ar rcs $@ $^

# Rule to compile runonnx.o
onnx/runonnx.o: onnx/runonnx.c onnx/runonnx.h
	$(CC) $(CFLAGS) -c $< -o $@
	@echo "Compiled runonnx.o"

# Rule to compile myfuncs.o
onnx/myfuncs.o: onnx/myfuncs.c onnx/myfuncs.h
	$(CC) $(CFLAGS) -c $< -o $@
	@echo "Compiled myfuncs.o"

# Rule to run the Go program
run: $(LIBS)
	go run main.go
	@echo "Go program executed"

# Clean up generated files
clean:
	rm -f $(OBJECTS) $(LIBS)
	@echo "Cleaned up"