.PHONY: generate

generate:
	@echo "Generating tags (fetching HTML specs from wooorm repository)..."
	go run ./cmd/gen-tags
	@echo "Running goimports..."
	goimports -w .
	@echo "Running gofmt..."
	gofmt -w .
	@echo "Generating SVG tags..."
	go run ./cmd/gen-svg
	@echo "Running goimports..."
	goimports -w .
	@echo "Running gofmt..."
	gofmt -w .
	@echo "Done!"
