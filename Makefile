# =============================================================================================== #
# DEVELOPMENT
# =============================================================================================== #

## run/backend: run the backend application
.PHONY: run/backend
run/backend:
	go run ./backend/cmd/api

## run/frontend: run the frontend application
.PHONY: run/frontend
run/frontend:
	go run ./frontend/cmd/api

## run: run the application
.PHONY: run
run: run/backend, run/frontend