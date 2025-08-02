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
	@go install github.com/vektra/mockery/v3@v3.5.0


## watch/api: hot reload the api server
.PHONY: watch/api
watch/api:
	air \
	--tmp_dir ".tmp" \
	--build.cmd "go build -o ./.tmp/api cmd/api/main.go" \
	--build.bin "./.tmp/api" \
	--build.delay "1000" \
	--build.exclude_dir ".tmp" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit false


## watch/job: hot reload the job server
.PHONY: watch/cron
watch/cron:
	air \
	--tmp_dir ".tmp" \
	--build.cmd "go build -o ./.tmp/cron ./cmd/cron/main.go" \
	--build.bin "./.tmp/cron" \
	--build.delay "1000" \
	--build.exclude_dir ".tmp" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit false


# start watch all processes in parallel.
.PHONY: start
start:
	make -j2 watch/job watch/api


## mock: generate the mocks
.PHONY: mock
mock:
	mockery


## test: run unit tests
.PHONY: test
test:
	ENV=.env.test go test -v ./...

