package svg

import (
	"fmt"
	"strings"

	"github.com/plainkit/html/cmd/gen/internal/files"
	"github.com/plainkit/html/cmd/gen/internal/svg/generator"
	"github.com/plainkit/html/cmd/gen/internal/svg/spec"
)

func Generate(outDir string) error {
	manager := files.NewManager(outDir)
	if err := manager.EnsureDir(); err != nil {
		return fmt.Errorf("ensure output directory: %w", err)
	}

	if err := manager.Clean(shouldRemoveGeneratedSVG); err != nil {
		return fmt.Errorf("clean generated files: %w", err)
	}

	specLoader := spec.NewLoader()
	attributesGenerator := generator.NewAttributesGenerator()
	tagGenerator := generator.NewTagGenerator()

	allSpecs, err := specLoader.LoadAllTagSpecs()
	if err != nil {
		return fmt.Errorf("load tag specs: %w", err)
	}

	fmt.Println("Generating SVG attributes file (excluding HTML ones)...")

	if err := generateAttributes(manager, specLoader, attributesGenerator, allSpecs); err != nil {
		return fmt.Errorf("generate SVG attributes: %w", err)
	}

	fmt.Println("Generating SVG tag files...")

	for _, tagSpec := range allSpecs {
		fileName := fmt.Sprintf("svg_%s.go", tagSpec.Name)
		source := tagGenerator.GenerateSource(tagSpec)

		if err := manager.WriteGoFile(fileName, source); err != nil {
			return fmt.Errorf("generate %s: %w", fileName, err)
		}
	}

	fmt.Printf("✅ SVG generation done (%d tags)\n", len(allSpecs))

	return nil
}

func shouldRemoveGeneratedSVG(name string) bool {
	return name == "svg_attrs.go" || (strings.HasPrefix(name, "svg_") && strings.HasSuffix(name, ".go"))
}

func generateAttributes(manager files.Manager, specLoader *spec.Loader, attrGen *generator.AttributesGenerator, allSpecs []spec.TagSpec) error {
	allAttributes := specLoader.CollectAllAttributes(allSpecs)
	fmt.Printf("Collected %d unique SVG attributes from all specs\n", len(allAttributes))

	source := attrGen.GenerateSource(allAttributes)

	return manager.WriteGoFile("svg_attrs.go", source)
}
