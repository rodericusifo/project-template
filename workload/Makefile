# COLOR VARIABLE
GREEN=\033[0;32m
RED=\033[0;31m
BLUE=\033[0;34m
LIGHT_BLUE=\033[1;34m
ORANGE=\033[0;33m
NOCOLOR=\033[0m

# STYLE VARIABLE
BLUE_TRIPLE_EQUALS=$(LIGHT_BLUE)===$(NOCOLOR)
GREEN_TRIPLE_RIGHT_ARROW=$(GREEN)>>>$(NOCOLOR)
GREEN_TRIPLE_LEFT_ARROW=$(GREEN)<<<$(NOCOLOR)

# FUNCTION
define log_action
$(BLUE_TRIPLE_EQUALS) $(ORANGE)$(1)$(NOCOLOR) $(BLUE_TRIPLE_EQUALS)
endef
define log_action_end
$(GREEN_TRIPLE_RIGHT_ARROW) $(1) $(GREEN_TRIPLE_LEFT_ARROW)
endef

# ENV VARIABLE
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
    ECHO_CMD = echo
else ifeq ($(UNAME_S),Darwin)
    ECHO_CMD = echo
else
    ECHO_CMD = echo -e
endif

GO_BIN_DIR := $(shell go env GOPATH 2>/dev/null)/bin
export PATH := $(GO_BIN_DIR):$(PATH)

.PHONY: gen-certs
gen-certs:
	@$(ECHO_CMD) "$(call log_action,Generate Certs Services ($(NAME) | $(ENV)))"
	chmod +x "./gen-certs.sh"
	bash ./gen-certs.sh
	@$(ECHO_CMD) "$(call log_action_end,All certificates generated successfully)"

.PHONY: check-proto-tools gen-proto
check-proto-tools:
	@command -v protoc >/dev/null || { echo "ERROR: protoc is required to generate protobuf files."; exit 1; }
	@command -v protoc-gen-go >/dev/null || { echo "ERROR: protoc-gen-go is required. Install it with: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; exit 1; }
	@command -v protoc-gen-go-grpc >/dev/null || { echo "ERROR: protoc-gen-go-grpc is required. Install it with: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; exit 1; }

gen-proto: check-proto-tools
	@$(ECHO_CMD) "$(call log_action,Generate Proto Services ($(NAME) | $(ENV)))"
	@for workload in $(WORKLOAD); do \
		make -C $${workload} gen-proto; \
	done
	@$(ECHO_CMD) "$(call log_action_end,All proto services generated successfully)"

.PHONY: run-workload
run-workload: gen-certs gen-proto
	@$(ECHO_CMD) "$(call log_action,Start App ($(NAME) | $(ENV)))"
	@for workload in $(WORKLOAD); do \
		$(ECHO_CMD) "$(call log_action,Start App [$${workload}] ($(NAME) | $(ENV)))"; \
		WORKLOAD=$${workload} ENV=$(ENV) SLUG=$(SLUG) VERSION=$$(cat $${workload}/VERSION) docker compose -p $(NAME)_workload_$${workload} --env-file $${workload}/env/$(ENV).env -f $${workload}/docker-compose.yaml up --build -d; \
	done
	@$(ECHO_CMD) "$(call log_action_end,App started successfully)"

.PHONY: stop-workload
stop-workload:
	@$(ECHO_CMD) "$(call log_action,Stop App ($(NAME) | $(ENV)))"
	@for workload in $(WORKLOAD); do \
		$(ECHO_CMD) "$(call log_action,Stop App [$${workload}] ($(NAME) | $(ENV)))"; \
		WORKLOAD=$${workload} ENV=$(ENV) SLUG=$(SLUG) VERSION=$(cat "$${workload}/VERSION") docker compose -p $(NAME)_workload_$${workload} --env-file $${workload}/env/$(ENV).env -f $${workload}/docker-compose.yaml down; \
	done
	@$(ECHO_CMD) "$(call log_action_end,App stopped successfully)"