all: help

.PHONY: help
help: Makefile
	@echo
	@echo "Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


## setup: install needed packages
.PHONY: setup
setup:
	@go install github.com/air-verse/air@latest
	@go install github.com/vektra/mockery/v2@v2.50.0


## start: build and run local project
.PHONY: start
start:
	air
	

## mock: generate the mocks
.PHONY: mock
mock:
	mockery --all


## test: run unit tests
.PHONY: test
test:
	ENV=.env.test go test -v ./...

