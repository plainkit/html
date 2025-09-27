package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/plainkit/html/cmd/gen/internal/html"
	"github.com/plainkit/html/cmd/gen/internal/svg"
)

func main() {
	outDir := flag.String("out", ".", "output directory for generated files")
	flag.Parse()

	if err := run(*outDir); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(outDir string) error {
	fmt.Println("Generating HTML...")
	if err := html.Generate(outDir); err != nil {
		return fmt.Errorf("html generation: %w", err)
	}

	fmt.Println("Generating SVG...")
	if err := svg.Generate(outDir); err != nil {
		return fmt.Errorf("svg generation: %w", err)
	}

	fmt.Println("✅ Generation complete")
	return nil
}
