.PHONY: update-specs

update-specs:
	@echo "Note: HTML specs are now fetched directly from wooorm repository"
	@echo "No local spec files need to be updated for HTML generation"
	@echo "Downloading SVG element attributes spec..."
	@mkdir -p svg/specs
	curl -s -L -o svg/specs/svg-element-attributes.js https://raw.githubusercontent.com/wooorm/svg-element-attributes/refs/heads/main/index.js
	@echo "SVG specs updated successfully!"


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
