.PHONY: run_lsp tests tests-json failed-tests create_test_project

export MY_TEST_PROJECT_DIRECTORY := $(shell sh ./create_test_project.sh)
run_lsp:
	@echo "Running LSP server..."
	@air -c ./config/air.toml 
create_test_project:
	@echo "creating test project"
tests: create_test_project
	@echo "Running tests..."
	@go test ./... -v -race -shuffle=on 
tests-json: create_test_project
	@echo "Running tests..."
	@go test ./... -v -race -shuffle=on -json 
failed-tests: create_test_project
	@echo "Running tests..."
	@go test ./... -v -race -shuffle=on -json | jq '.|select(.Action=="fail" and .Test!=null)'

