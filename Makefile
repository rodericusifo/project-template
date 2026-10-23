PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
include $(PROJECT_ROOT)/stack.mk
include $(PROJECT_ROOT)/colors.mk

define check_env
	@if [ -z "$(ENV)" ]; then \
		echo "ERROR: ENV is required. Usage: make $(1) ENV=dev"; \
		exit 1; \
	fi
endef

# ENV VARIABLE
NAME := $(notdir $(CURDIR))
SLUG := $(subst -,_,$(NAME))

.PHONY: setup-hooks
setup-hooks:
	@$(PRINTF_CMD) "$(call log_action,Setting up Git Hooks)"
	chmod +x "./script/setup-hooks.sh"
	bash ./script/setup-hooks.sh
	@$(PRINTF_CMD) "$(call log_action_end,Git Hooks setup completed)"

.PHONY: setup-env
setup-env:
	$(call check_env,$(MAKECMDGOALS))
	@$(PRINTF_CMD) "$(call log_action,Setting up Environment Files)"
	chmod +x "./script/setup-env.sh"
	bash ./script/setup-env.sh $(ENV)
	@$(PRINTF_CMD) "$(call log_action_end,Environment files setup completed)"

.PHONY: start
start:
	$(call check_env,$(MAKECMDGOALS))
	@$(PRINTF_CMD) "$(call log_action,Generate App Network ($(NAME) | $(ENV)))"
	docker network ls | grep $(SLUG)_$(ENV) || docker network create $(SLUG)_$(ENV)
	@make -C infrastructure/database/sql DATABASE_SQL="$(DATABASE_SQL)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" run-database-sql
	@make -C infrastructure/cache CACHE="$(CACHE)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" run-cache
	@make -C infrastructure/messaging MESSAGING="$(MESSAGING)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" run-messaging
	@make -C workload WORKLOAD="$(WORKLOAD)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" run-workload

.PHONY: stop
stop:
	$(call check_env,$(MAKECMDGOALS))
	@make -C infrastructure/database/sql DATABASE_SQL="$(DATABASE_SQL)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" stop-database-sql
	@make -C infrastructure/cache CACHE="$(CACHE)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" stop-cache
	@make -C infrastructure/messaging MESSAGING="$(MESSAGING)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" stop-messaging
	@make -C workload WORKLOAD="$(WORKLOAD)" SLUG="$(SLUG)" NAME="$(NAME)" ENV="$(ENV)" stop-workload