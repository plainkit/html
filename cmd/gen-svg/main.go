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

	fmt.Println("Generating SVG tag files...")
	for _, tagSpec := range allSpecs {
		fileName := fmt.Sprintf("tag_%s.go", tagSpec.Name)
		source := tagGenerator.GenerateSource(tagSpec)

		if err := fileManager.WriteFormattedFile(fileName, source); err != nil {
			return fmt.Errorf("generate %s: %w", fileName, err)
		}
	}

	fmt.Printf("✅ Successfully generated %d SVG tag files\n", len(allSpecs))
	return nil
}
