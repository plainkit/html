.PHONY: generate

generate:
	@echo "Generating HTML tags from gostar..."
	cd cmd/gen-tags && go run main.go -out ../..
	@echo "Generating SVG tags..."
	go run ./cmd/gen-svg
	@echo "Running goimports..."
	goimports -w .
	@echo "Running gofmt..."
	gofmt -w .
	@echo "Done!"
