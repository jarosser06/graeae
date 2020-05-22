COMMAND_NAME = graeae
BUILDS_DIR = bin

ifeq ($(OS), Windows_NT)
	BINARY_NAMWE = $(COMMAND_NAME).exe
else
	BINARY_NAME = $(COMMAND_NAME)
endif

all: build

build: ${BUILDS_DIR}/${BINARY_NAME}
${BUILDS_DIR}/${BINARY_NAME}:
	@echo "Building binary..."
	go build -o ${BUILDS_DIR}/${BINARY_NAME}

clean:
	@echo "Cleaning up..."
	rm -rf bin

rebuild: clean build

cross_compile: linux darwin windows

linux: ${BUILDS_DIR}/linux/${COMMAND_NAME}
${BUILDS_DIR}/linux/${COMMAND_NAME}:
	@mkdir -p ${BUILDS_DIR}/linux
	GOOS=linux go build -v -o bin/linux/${COMMAND_NAME}

darwin: ${BUILDS_DIR}/darwin/${COMMAND_NAME}
${BUILDS_DIR}/darwin/${COMMAND_NAME}:
	@mkdir -p ${BUILDS_DIR}/darwin
	GOOS=darwin go build -v -o bin/darwin/${COMMAND_NAME}

windows: ${BUILDS_DIR}/windows/${COMMAND_NAME}.exe
${BUILDS_DIR}/windows/${COMMAND_NAME}.exe:
	@mkdir -p ${BUILDS_DIR}/windows
	GOOS=windows go build -v -o bin/windows/${COMMAND_NAME}.exe

.PHONY: linux darwin windows cross_compile rebuild clean
