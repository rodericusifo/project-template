PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))..)
include $(PROJECT_ROOT)/colors.mk

# ENV VARIABLE

GO_BIN_DIR := $(shell go env GOPATH 2>/dev/null)/bin
export PATH := $(GO_BIN_DIR):$(PATH)

.PHONY: gen-certs
gen-certs:
	@$(PRINTF_CMD) "$(call log_action,Generate Certs Services ($(NAME) | $(ENV)))"
	chmod +x "./gen-certs.sh"
	bash ./gen-certs.sh
	@$(PRINTF_CMD) "$(call log_action_end,All certificates generated successfully)"

.PHONY: check-proto-tools gen-proto
check-proto-tools:
	@command -v protoc >/dev/null || { echo "ERROR: protoc is required to generate protobuf files."; exit 1; }
	@command -v protoc-gen-go >/dev/null || { echo "ERROR: protoc-gen-go is required. Install it with: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; exit 1; }
	@command -v protoc-gen-go-grpc >/dev/null || { echo "ERROR: protoc-gen-go-grpc is required. Install it with: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; exit 1; }

gen-proto: check-proto-tools
	@$(PRINTF_CMD) "$(call log_action,Generate Proto Services ($(NAME) | $(ENV)))"
	@for workload in $(WORKLOAD); do \
		make -C $${workload} gen-proto; \
	done
	@$(PRINTF_CMD) "$(call log_action_end,All proto services generated successfully)"

.PHONY: run-workload
run-workload: gen-certs gen-proto
	@$(PRINTF_CMD) "$(call log_action,Start App ($(NAME) | $(ENV)))"
	@for workload in $(WORKLOAD); do \
		$(PRINTF_CMD) "$(call log_action,Start App [$${workload}] ($(NAME) | $(ENV)))"; \
		WORKLOAD=$${workload} ENV=$(ENV) SLUG=$(SLUG) VERSION=$$(cat $${workload}/VERSION) docker compose -p $(NAME)_workload_$${workload} --env-file $${workload}/env/$(ENV).env -f $${workload}/docker-compose.yaml up --build -d; \
	done
	@$(PRINTF_CMD) "$(call log_action_end,App started successfully)"

.PHONY: stop-workload
stop-workload:
	@$(PRINTF_CMD) "$(call log_action,Stop App ($(NAME) | $(ENV)))"
	@for workload in $(WORKLOAD); do \
		$(PRINTF_CMD) "$(call log_action,Stop App [$${workload}] ($(NAME) | $(ENV)))"; \
		WORKLOAD=$${workload} ENV=$(ENV) SLUG=$(SLUG) VERSION=$(cat "$${workload}/VERSION") docker compose -p $(NAME)_workload_$${workload} --env-file $${workload}/env/$(ENV).env -f $${workload}/docker-compose.yaml down; \
	done
	@$(PRINTF_CMD) "$(call log_action_end,App stopped successfully)"