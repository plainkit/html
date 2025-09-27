package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/plainkit/html/cmd/gen-svg/internal/files"
	"github.com/plainkit/html/cmd/gen-svg/internal/generator"
	"github.com/plainkit/html/cmd/gen-svg/internal/spec"
)

func main() {
	outDir := flag.String("out", "svg", "output directory for generated tags")
	flag.Parse()

	if err := run(*outDir); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(outDir string) error {
	fileManager := files.NewManager(outDir)
	specLoader := spec.NewLoader()
	tagGenerator := generator.NewTagGenerator()
	globalGenerator := generator.NewGlobalGenerator()
	attributesGenerator := generator.NewAttributesGenerator()

	if err := fileManager.EnsureOutputDir(); err != nil {
		return fmt.Errorf("ensure output directory: %w", err)
	}

	if err := fileManager.CleanGeneratedFiles(); err != nil {
		return fmt.Errorf("clean generated files: %w", err)
	}

	allSpecs, err := specLoader.LoadAllTagSpecs()
	if err != nil {
		return fmt.Errorf("load tag specs: %w", err)
	}

	fmt.Println("Generating SVG global attributes...")
	if err := generateSvgGlobal(specLoader, fileManager, globalGenerator); err != nil {
		return fmt.Errorf("generate SVG global: %w", err)
	}

	fmt.Println("Generating SVG centralized attributes file...")
	if err := generateSvgAttributes(specLoader, fileManager, attributesGenerator, allSpecs); err != nil {
		return fmt.Errorf("generate SVG attributes: %w", err)
	}

	fmt.Println("Generating SVG tag files...")
	for _, tagSpec := range allSpecs {
		fileName := fmt.Sprintf("tag_%s.go", tagSpec.Name)
		source := tagGenerator.GenerateSource(tagSpec)

		if err := fileManager.WriteFormattedFile(fileName, source); err != nil {
			return fmt.Errorf("generate %s: %w", fileName, err)
		}
	}

	fmt.Printf("✅ Successfully generated %d SVG tag files and supporting code\n", len(allSpecs))
	return nil
}

func generateSvgGlobal(specLoader *spec.Loader, fileManager *files.Manager, globalGen *generator.GlobalGenerator) error {
	globalAttrs, err := specLoader.LoadGlobalAttributes()
	if err != nil {
		return fmt.Errorf("load global attributes: %w", err)
	}

	source := globalGen.GenerateSource(globalAttrs)
	return fileManager.WriteFormattedFile("svg_global.go", source)
}

func generateSvgAttributes(specLoader *spec.Loader, fileManager *files.Manager, attrGen *generator.AttributesGenerator, allSpecs []spec.TagSpec) error {
	allAttributes := specLoader.CollectAllAttributes(allSpecs)
	fmt.Printf("Collected %d unique SVG attributes from all specs\n", len(allAttributes))

	source := attrGen.GenerateSource(allAttributes)
	return fileManager.WriteFormattedFile("svg_attrs.go", source)
}
