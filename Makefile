ifneq (,$(wildcard .env))
    include .env
    export
endif

lint:
	sh scripts/lint.sh
