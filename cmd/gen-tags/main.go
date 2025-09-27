package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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
	fileManager := files.NewManager(outDir)
	specLoader := spec.NewLoader(specsDir)

	if err := fileManager.EnsureOutputDir(); err != nil {
		return fmt.Errorf("ensure output directory: %w", err)
	}

	if err := fileManager.CleanGeneratedFiles(); err != nil {
		return fmt.Errorf("clean generated files: %w", err)
	}

	fmt.Println("Loading HTML specifications from gostar...")
	allSpecs, err := specLoader.LoadAllTagSpecsFromGostar()
	if err != nil {
		return fmt.Errorf("load tag specs: %w", err)
	}
	fmt.Printf("Loaded %d tag specifications\n", len(allSpecs))

	fmt.Println("Generating centralized attributes file...")
	allAttributes, err := generateAttributesFile(specLoader, fileManager, allSpecs)
	if err != nil {
		return fmt.Errorf("generate attributes file: %w", err)
	}

	fmt.Println("Generating core_global.go...")
	if err := generateCoreGlobal(specLoader, fileManager, allAttributes); err != nil {
		return fmt.Errorf("generate core global: %w", err)
	}

	fmt.Println("Generating tag files...")
	if err := generateTagFiles(fileManager, allSpecs); err != nil {
		return fmt.Errorf("generate tag files: %w", err)
	}

	fmt.Println("Generating core_node.go with all Node apply methods...")
	if err := generateCoreNodeFile(fileManager, allSpecs); err != nil {
		return fmt.Errorf("generate core node file: %w", err)
	}

	fmt.Printf("✅ Successfully generated %d tag files and supporting code\n", len(allSpecs))
	return nil
}

func generateCoreGlobal(specLoader *spec.Loader, fileManager *files.Manager, centralizedAttrs map[string]spec.Attribute) error {
	globalAttrs, err := specLoader.LoadGlobalAttributesFromGostar()
	if err != nil {
		return fmt.Errorf("load global attributes: %w", err)
	}

	// Filter out global attributes that are already in the centralized attrs file
	filteredGlobalAttrs := make([]spec.Attribute, 0)
	for _, globalAttr := range globalAttrs {
		key := strings.ToLower(globalAttr.Attr)
		if _, exists := centralizedAttrs[key]; !exists {
			filteredGlobalAttrs = append(filteredGlobalAttrs, globalAttr)
		}
	}

	globalGen := generator.NewGlobalGenerator()
	source := globalGen.GenerateSource(filteredGlobalAttrs)

	return fileManager.WriteFormattedFile("core_global.go", source)
}

func generateAttributesFile(specLoader *spec.Loader, fileManager *files.Manager, allSpecs []spec.TagSpec) (map[string]spec.Attribute, error) {
	allAttributes := specLoader.CollectAllAttributes(allSpecs)
	fmt.Printf("Collected %d unique attributes from all specs\n", len(allAttributes))

	attrGen := generator.NewAttributesGenerator()
	source := attrGen.GenerateSource(allAttributes)

	err := fileManager.WriteFormattedFile("attrs.go", source)
	return allAttributes, err
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
