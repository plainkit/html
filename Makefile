.PHONY: generate

generate:
	@echo "Generating HTML tags from gostar..."
	cd cmd/gen-tags && go run main.go -out ../..
	@echo "Generating SVG tags..."
	cd cmd/gen-svg && go run main.go -out ../../svg
	@echo "Running goimports..."
	goimports -w .
	@echo "Running gofmt..."
	gofmt -w .
	@echo "Done!"
