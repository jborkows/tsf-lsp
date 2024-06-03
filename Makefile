.PHONY: run_lsp tests tests-json failed-tests
run_lsp:
	@echo "Running LSP server..."
	@air -c ./config/air.toml 
tests:
	@echo "Running tests..."
	@go test ./... -v -race -shuffle=on 
tests-json:
	@echo "Running tests..."
	@go test ./... -v -race -shuffle=on -json 

failed-tests:
	@echo "Running tests..."
	@go test ./... -v -race -shuffle=on -json | jq '.|select(.Action=="fail" and .Test!=null)'

