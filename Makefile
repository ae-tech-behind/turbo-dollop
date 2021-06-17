.PHONY: build
build: dep ## Build
	@echo "::: building app"
	go build -o ./bin/app
	@echo "::: done"

.PHONY: run
run: build ## Run application
	@echo "::: running app"
	./bin/app

.PHONY: dep
dep: ## Download dependecies
	@echo "::: downloading dependencies"
	go mod tidy
	@echo "::: done"

.PHONY: test
test: ## Unit test
	@echo "::: excecuting unit-test"
	go test ./...
	@echo "::: done"

