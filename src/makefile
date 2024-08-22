# Compiler and flags
CC = gcc
CFLAGS = -arch arm64 -I/opt/homebrew/Cellar/onnxruntime/1.17.1/include/onnxruntime

# Target libraries and object files
LIBS = librunonnx.a libmyfuncs.a
OBJECTS = runonnx.o myfuncs.o

# Default target: build libraries and run the Go program
all: $(LIBS) run

# Rule to create librunonnx.a
librunonnx.a: runonnx.o
	ar rcs $@ $^

# Rule to create libmyfuncs.a
libmyfuncs.a: myfuncs.o
	ar rcs $@ $^

# Rule to compile runonnx.o
runonnx.o: runonnx.c
	$(CC) $(CFLAGS) -c $< -o $@

# Rule to compile myfuncs.o
myfuncs.o: myfuncs.c
	$(CC) $(CFLAGS) -c $< -o $@

# Rule to run the Go program
run: $(LIBS)
	go run main.go

# Clean up generated files
clean:
	rm -f $(OBJECTS) $(LIBS)