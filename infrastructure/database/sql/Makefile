SHELL := /bin/bash
PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))../../..)
include $(PROJECT_ROOT)/colors.mk

# ENV VARIABLE
AVAILABLE_DATABASE_SQL := postgres

.PHONY: check-database-sql
check-database-sql:
	@$(PRINTF_CMD) "$(call log_action,Check App Databases SQL ($(NAME) | $(ENV)))"
	@for database_sql in $(DATABASE_SQL); do \
		if [[ ! " $(AVAILABLE_DATABASE_SQL) " =~ " $${database_sql} " ]]; then \
			$(PRINTF_CMD) "$(RED)ERROR: Database SQL [$${database_sql}] is not available. Available Database SQL are: $(AVAILABLE_DATABASE_SQL)$(NOCOLOR)"; \
			exit 1; \
		fi; \
	done
	@$(PRINTF_CMD) "$(call log_action_end,All Databases SQL are available)"

.PHONY: run-database-sql
run-database-sql: check-database-sql
	@$(PRINTF_CMD) "$(call log_action,Start App Databases SQL ($(NAME) | $(ENV)))"
	@for database_sql in $(DATABASE_SQL); do \
		$(PRINTF_CMD) "$(call log_action,Generate App Database SQL Volumes [$${database_sql}] ($(NAME) | $(ENV)))"; \
		if [ "$${database_sql}" == "postgres" ]; then \
			chmod +x "postgres/primary/initdb/db.sh"; \
			chmod +x "postgres/primary/initdb/role.sh"; \
		fi; \
		docker volume ls | grep $(SLUG)_primary_$${database_sql}_data_$(ENV) || docker volume create --name $(SLUG)_primary_$${database_sql}_data_$(ENV); \
		docker volume ls | grep $(SLUG)_replica_$${database_sql}_data_$(ENV) || docker volume create --name $(SLUG)_replica_$${database_sql}_data_$(ENV); \
		$(PRINTF_CMD) "$(call log_action,Start App Database SQL [$${database_sql}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_database_sql --env-file $${database_sql}/env/$(ENV).env -f $${database_sql}/docker-compose.yaml up --build -d; \
	done
	@$(PRINTF_CMD) "$(call log_action_end,Databases SQL started successfully)"

.PHONY: stop-database-sql
stop-database-sql: check-database-sql
	@$(PRINTF_CMD) "$(call log_action,Stop App Databases SQL ($(NAME) | $(ENV)))"
	@for database_sql in $(DATABASE_SQL); do \
		$(PRINTF_CMD) "$(call log_action,Stop App Database SQL [$${database_sql}] ($(NAME) | $(ENV)))"; \
		ENV=$(ENV) SLUG=$(SLUG) docker compose -p $(NAME)_database_sql --env-file $${database_sql}/env/$(ENV).env -f $${database_sql}/docker-compose.yaml down; \
	done
	@$(PRINTF_CMD) "$(call log_action_end,Databases SQL stopped successfully)"