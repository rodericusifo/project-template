GO_BIN_DIR := $(shell go env GOPATH 2>/dev/null)/bin
export PATH := $(GO_BIN_DIR):$(PATH)

.PHONY: check-proto-tools
check-proto-tools:
	@command -v protoc >/dev/null || { echo "ERROR: protoc is required to generate protobuf files."; exit 1; }
	@command -v protoc-gen-go >/dev/null || { echo "ERROR: protoc-gen-go is required. Install it with: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; exit 1; }
	@command -v protoc-gen-go-grpc >/dev/null || { echo "ERROR: protoc-gen-go-grpc is required. Install it with: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; exit 1; }

.PHONY: proto-auth-service
proto-auth-service:
	protoc \
		--proto_path=external/client/auth-service/http/grpc/proto \
		--proto_path=/usr/include \
		--go_out=external/client/auth-service/http/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=external/client/auth-service/http/grpc/proto --go-grpc_opt=paths=source_relative \
		external/client/auth-service/http/grpc/proto/auth/auth.proto \
		external/client/auth-service/http/grpc/proto/user/user.proto
.PHONY: proto-author-service
proto-author-service:
	protoc \
		--proto_path=external/client/author-service/http/grpc/proto \
		--proto_path=/usr/include \
		--go_out=external/client/author-service/http/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=external/client/author-service/http/grpc/proto --go-grpc_opt=paths=source_relative \
		external/client/author-service/http/grpc/proto/author/author.proto
.PHONY: proto-category-service
proto-category-service:
	protoc \
		--proto_path=external/client/category-service/http/grpc/proto \
		--proto_path=/usr/include \
		--go_out=external/client/category-service/http/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=external/client/category-service/http/grpc/proto --go-grpc_opt=paths=source_relative \
		external/client/category-service/http/grpc/proto/category/category.proto
.PHONY: proto-book-service
proto-book-service:
	protoc \
		--proto_path=external/client/book-service/http/grpc/proto \
		--proto_path=/usr/include \
		--go_out=external/client/book-service/http/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=external/client/book-service/http/grpc/proto --go-grpc_opt=paths=source_relative \
		external/client/book-service/http/grpc/proto/book/book.proto

.PHONY: gen-proto
gen-proto: check-proto-tools proto-auth-service proto-author-service proto-category-service proto-book-service