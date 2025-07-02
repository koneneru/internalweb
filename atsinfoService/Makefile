include .envrc

# =============================================================================================== #
# HELPERS
# =============================================================================================== #

## help: print this help information
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' | sort


# =============================================================================================== #
# DEVELOPMENT
# =============================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api -dbd-dsn=${DIRECTUM_DB_DSN} -dbs-dsn=${SVET_DB_DSN}