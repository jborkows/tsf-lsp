.PHONY: run_lsp tests
run_lsp:
	@echo "Running LSP server..."
	@air -c ./config/air.toml 
tests:
	@echo "Running tests..."
	@go test ./... -v -race -shuffle=on 

