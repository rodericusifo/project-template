NO_COLOR ?= 0
FORCE_COLOR ?= 0

ifeq ($(NO_COLOR),1)
GREEN :=
RED :=
BLUE :=
LIGHT_BLUE :=
ORANGE :=
NOCOLOR :=
else ifeq ($(FORCE_COLOR),1)
ESC := $(shell printf '\033')
GREEN := $(ESC)[0;32m
RED := $(ESC)[0;31m
BLUE := $(ESC)[0;34m
LIGHT_BLUE := $(ESC)[1;34m
ORANGE := $(ESC)[0;33m
NOCOLOR := $(ESC)[0m
else ifneq ($(TERM),dumb)
ESC := $(shell printf '\033')
GREEN := $(ESC)[0;32m
RED := $(ESC)[0;31m
BLUE := $(ESC)[0;34m
LIGHT_BLUE := $(ESC)[1;34m
ORANGE := $(ESC)[0;33m
NOCOLOR := $(ESC)[0m
endif

PRINTF_CMD := printf '%s\n'

BLUE_TRIPLE_EQUALS := $(LIGHT_BLUE)===$(NOCOLOR)
GREEN_TRIPLE_RIGHT_ARROW := $(GREEN)>>>$(NOCOLOR)
GREEN_TRIPLE_LEFT_ARROW := $(GREEN)<<<$(NOCOLOR)

define log_action
$(BLUE_TRIPLE_EQUALS) $(ORANGE)$(1)$(NOCOLOR) $(BLUE_TRIPLE_EQUALS)
endef

define log_action_end
$(GREEN_TRIPLE_RIGHT_ARROW) $(1) $(GREEN_TRIPLE_LEFT_ARROW)
endef
