# Compiler
CC = gcc

# Static libraries and object files
LIBS = src/onnx/librunonnx.a src/onnx/libmyfuncs.a
OBJECTS = src/onnx/runonnx.o src/onnx/myfuncs.o

# Default target: build libraries and run the Go program
all: $(LIBS) run

# Rule to create librunonnx.a
src/onnx/librunonnx.a: src/onnx/runonnx.o
	ar rcs $@ $^

# Rule to create libmyfuncs.a
src/onnx/libmyfuncs.a: src/onnx/myfuncs.o
	ar rcs $@ $^

# Rule to compile runonnx.o
src/onnx/runonnx.o: src/onnx/runonnx.c src/onnx/runonnx.h
	$(CC) $(CFLAGS) -c $< -o $@
	@echo "Compiled runonnx.o"

# Rule to compile myfuncs.o
src/onnx/myfuncs.o: src/onnx/myfuncs.c src/onnx/myfuncs.h
	$(CC) $(CFLAGS) -c $< -o $@
	@echo "Compiled myfuncs.o"

# Rule to run the Go program
run: $(LIBS)
	cd src && go run main.go
	@echo "Go program executed"

# Clean up generated files
clean:
	rm -f $(OBJECTS) $(LIBS)
	@echo "Cleaned up"