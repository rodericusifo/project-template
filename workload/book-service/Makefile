GO_BIN_DIR := $(shell go env GOPATH 2>/dev/null)/bin
export PATH := $(GO_BIN_DIR):$(PATH)

.PHONY: gen-proto
.PHONY: check-proto-tools
check-proto-tools:
	@command -v protoc >/dev/null || { echo "ERROR: protoc is required to generate protobuf files."; exit 1; }
	@command -v protoc-gen-go >/dev/null || { echo "ERROR: protoc-gen-go is required. Install it with: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; exit 1; }
	@command -v protoc-gen-go-grpc >/dev/null || { echo "ERROR: protoc-gen-go-grpc is required. Install it with: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; exit 1; }

gen-proto: check-proto-tools
	protoc \
	--proto_path=internal/proto \
	--proto_path=/usr/include \
    --go_out=internal/proto --go_opt=paths=source_relative \
    --go-grpc_out=internal/proto --go-grpc_opt=paths=source_relative \
    internal/proto/book/book.proto