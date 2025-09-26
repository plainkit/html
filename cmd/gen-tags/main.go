package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/plainkit/html/cmd/gen-tags/internal/files"
	"github.com/plainkit/html/cmd/gen-tags/internal/generator"
	"github.com/plainkit/html/cmd/gen-tags/internal/spec"
)

func main() {
	specsDir := flag.String("specs", "specs", "path to specs directory")
	outDir := flag.String("out", ".", "output directory for generated tags")
	flag.Parse()

	if err := run(*specsDir, *outDir); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(specsDir, outDir string) error {
	// Initialize components
	fileManager := files.NewManager(outDir)
	specLoader := spec.NewLoader(specsDir)

	// Ensure output directory exists
	if err := fileManager.EnsureOutputDir(); err != nil {
		return fmt.Errorf("ensure output directory: %w", err)
	}

	// Clean previously generated files
	if err := fileManager.CleanGeneratedFiles(); err != nil {
		return fmt.Errorf("clean generated files: %w", err)
	}

	// Load all specifications from wooorm data
	fmt.Println("Fetching HTML specifications from wooorm repository...")
	allSpecs, err := specLoader.LoadAllTagSpecsFromWooorm()
	if err != nil {
		return fmt.Errorf("load tag specs: %w", err)
	}
	fmt.Printf("Loaded %d tag specifications\n", len(allSpecs))

	// Generate core global attributes file
	fmt.Println("Generating core_global.go...")
	if err := generateCoreGlobal(specLoader, fileManager); err != nil {
		return fmt.Errorf("generate core global: %w", err)
	}

	// Generate centralized attributes file
	fmt.Println("Generating centralized attributes file...")
	if err := generateAttributesFile(specLoader, fileManager, allSpecs); err != nil {
		return fmt.Errorf("generate attributes file: %w", err)
	}

	// Generate individual tag files
	fmt.Println("Generating tag files...")
	if err := generateTagFiles(fileManager, allSpecs); err != nil {
		return fmt.Errorf("generate tag files: %w", err)
	}

	// Generate complete core_node.go file
	fmt.Println("Generating core_node.go with all Node apply methods...")
	if err := generateCoreNodeFile(fileManager, allSpecs); err != nil {
		return fmt.Errorf("generate core node file: %w", err)
	}

	fmt.Printf("✅ Successfully generated %d tag files and supporting code\n", len(allSpecs))
	return nil
}

func generateCoreGlobal(specLoader *spec.Loader, fileManager *files.Manager) error {
	globalAttrs, err := specLoader.LoadGlobalAttributesFromWooorm()
	if err != nil {
		return fmt.Errorf("load global attributes: %w", err)
	}

	// Generate source code
	globalGen := generator.NewGlobalGenerator()
	source := globalGen.GenerateSource(globalAttrs)

	// Write to file
	return fileManager.WriteFormattedFile("core_global.go", source)
}

func generateAttributesFile(specLoader *spec.Loader, fileManager *files.Manager, allSpecs []spec.TagSpec) error {
	allAttributes := specLoader.CollectAllAttributes(allSpecs)
	fmt.Printf("Collected %d unique attributes from all specs\n", len(allAttributes))

	attrGen := generator.NewAttributesGenerator()
	source := attrGen.GenerateSource(allAttributes)

	return fileManager.WriteFormattedFile("attrs.go", source)
}

func generateTagFiles(fileManager *files.Manager, allSpecs []spec.TagSpec) error {
	tagGen := generator.NewTagGenerator()

	for _, tagSpec := range allSpecs {
		fileName := fmt.Sprintf("tag_%s.go", tagSpec.Name)
		source := tagGen.GenerateSource(tagSpec)

		if err := fileManager.WriteFormattedFile(fileName, source); err != nil {
			return fmt.Errorf("generate %s: %w", fileName, err)
		}
	}

	return nil
}

func generateCoreNodeFile(fileManager *files.Manager, allSpecs []spec.TagSpec) error {
	nodeGen := generator.NewNodeGenerator()
	content := nodeGen.GenerateCompleteFile(allSpecs)

	return fileManager.WriteFormattedFile("core_node.go", content)
}
