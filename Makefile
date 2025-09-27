.PHONY: generate

generate:
	@echo "Generating HTML and SVG from gostar..."
	cd cmd/gen && go run . -out ../..
	@echo "Running goimports..."
	goimports -w .
	@echo "Running gofmt..."
	gofmt -w .
	@echo "Done!"
