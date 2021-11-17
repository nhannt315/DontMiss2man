PROJ_DIR:=$(PWD)/..

export GOBIN:=$(PROJ_DIR)/bin
export PATH:=$(GOBIN):$(PATH)

OAPI_VERSION:=v1
OAPI_DIR:=$(PWD)/$(OAPI_VERSION)
OAPI_FILE:=$(OAPI_DIR)/real-estate-api.yml
OAPI_SRC_DIR:=$(PROJ_DIR)/internal/openapi/$(OAPI_VERSION)
OAPI_GO_PACKAGE=openapi

## openapiファイルからソースコードを生成
gen-oapi:
	$(GOBIN)/oapi-codegen -package $(OAPI_GO_PACKAGE) -generate spec -o $(OAPI_SRC_DIR)/real_estate_api_spec.gen.go  $(OAPI_FILE)
	$(GOBIN)/oapi-codegen -package $(OAPI_GO_PACKAGE) -generate types -o $(OAPI_SRC_DIR)/real_estate_api_types.gen.go  $(OAPI_FILE)
	$(GOBIN)/oapi-codegen -package $(OAPI_GO_PACKAGE) -generate server -o $(OAPI_SRC_DIR)/real_estate_api_server.gen.go  $(OAPI_FILE)
