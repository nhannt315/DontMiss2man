TARGET=dm2m_web
TARGET_BUILDER=$(TARGET)_builder
DOCKERFILE=./Dockerfile
ROOT_DIR=.

BUILD_PROLOGUE=DOCKER_BUILDKIT=1 docker build --progress=plain -f $(DOCKERFILE)

.PHONY: builder
builder:
	$(BUILD_PROLOGUE) -t $(TARGET_BUILDER) --target builder $(ROOT_DIR)

.PHONY: local
local:
	$(BUILD_PROLOGUE) -t $(TARGET) --build-arg BUILD_TARGET=local $(ROOT_DIR)

.PHONY: stg
stg:
	$(BUILD_PROLOGUE) -t $(TARGET) --build-arg BUILD_TARGET=stg $(ROOT_DIR)

.PHONY: prod
prod:
	$(BUILD_PROLOGUE) -t $(TARGET) --build-arg BUILD_TARGET=prod $(ROOT_DIR)

# misc

.PHONY: clean
clean:
	docker image rm $(TARGET) $(TARGET_BUILDER)