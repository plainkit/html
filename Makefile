.PHONY: update-specs

update-specs:
	@echo "Downloading MDN browser-compat-data..."
	@mkdir -p tmp
	curl -s -L -o tmp/bcd.zip https://github.com/mdn/browser-compat-data/archive/refs/heads/main.zip
	@echo "Extracting archive..."
	cd tmp && unzip -q bcd.zip
	@echo "Preparing specs directory..."
	mkdir -p specs
	rm -f specs/*
	@echo "Copying HTML JSON files..."
	find tmp/browser-compat-data-main/html -type f -name '*.json' -exec cp {} specs/ \;
	rm -f specs/text.json
	@echo "Cleaning up temporary files..."
	rm -rf tmp
	@echo "Downloading SVG element attributes spec..."
	@mkdir -p svg/specs
	curl -s -L -o svg/specs/svg-element-attributes.js https://raw.githubusercontent.com/wooorm/svg-element-attributes/refs/heads/main/index.js
	@echo "SVG specs updated successfully!"


generate:
	@echo "Generating tags..."
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
