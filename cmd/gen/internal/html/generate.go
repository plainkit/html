package html

import (
	"fmt"
	"strings"

	"github.com/plainkit/html/cmd/gen/internal/files"
	"github.com/plainkit/html/cmd/gen/internal/html/generator"
	"github.com/plainkit/html/cmd/gen/internal/html/spec"
)

func Generate(outDir string) error {
	manager := files.NewManager(outDir)
	if err := manager.EnsureDir(); err != nil {
		return fmt.Errorf("ensure output directory: %w", err)
	}

	if err := manager.Clean(shouldRemoveGeneratedHTML); err != nil {
		return fmt.Errorf("clean generated files: %w", err)
	}

	specLoader := spec.NewLoader()

	fmt.Println("Loading HTML element definitions from github.com/plainkit/tags...")

	allSpecs, err := specLoader.LoadAllTagSpecs()
	if err != nil {
		return fmt.Errorf("load tag specs: %w", err)
	}

	fmt.Printf("Loaded %d tag specifications\n", len(allSpecs))

	fmt.Println("Generating centralized attributes file...")

	allAttributes, err := generateAttributesFile(manager, allSpecs, specLoader)
	if err != nil {
		return fmt.Errorf("generate attributes file: %w", err)
	}

	fmt.Println("Generating core_global.go...")

	if err := generateCoreGlobal(manager, allAttributes, specLoader); err != nil {
		return fmt.Errorf("generate core global: %w", err)
	}

	fmt.Println("Generating tag files...")

	if err := generateTagFiles(manager, allSpecs); err != nil {
		return fmt.Errorf("generate tag files: %w", err)
	}

	fmt.Println("Generating core_node.go with all Node apply methods...")

	if err := generateCoreNodeFile(manager, allSpecs); err != nil {
		return fmt.Errorf("generate core node file: %w", err)
	}

	fmt.Printf("✅ HTML generation done (%d tags)\n", len(allSpecs))

	return nil
}

func shouldRemoveGeneratedHTML(name string) bool {
	return strings.HasPrefix(name, "tag_") || name == "attrs.go" || name == "core_options.go" || name == "core_global.go"
}

func generateCoreGlobal(manager files.Manager, centralizedAttrs map[string]spec.Attribute, specLoader *spec.Loader) error {
	globalAttrs, err := specLoader.LoadGlobalAttributes()
	if err != nil {
		return fmt.Errorf("load global attributes: %w", err)
	}

	filtered := make([]spec.Attribute, 0, len(globalAttrs))
	for _, attr := range globalAttrs {
		if _, exists := centralizedAttrs[strings.ToLower(attr.Attr)]; !exists {
			filtered = append(filtered, attr)
		}
	}

	source := generator.NewGlobalGenerator().GenerateSource(filtered)

	return manager.WriteGoFile("core_global.go", source)
}

func generateAttributesFile(manager files.Manager, allSpecs []spec.TagSpec, specLoader *spec.Loader) (map[string]spec.Attribute, error) {
	allAttributes := specLoader.CollectAllAttributes(allSpecs)
	fmt.Printf("Collected %d unique attributes from all specs\n", len(allAttributes))

	source := generator.NewAttributesGenerator().GenerateSource(allAttributes)
	if err := manager.WriteGoFile("attrs.go", source); err != nil {
		return nil, err
	}

	return allAttributes, nil
}

func generateTagFiles(manager files.Manager, allSpecs []spec.TagSpec) error {
	tagGen := generator.NewTagGenerator()

	for _, tagSpec := range allSpecs {
		fileName := fmt.Sprintf("tag_%s.go", tagSpec.Name)
		if tagSpec.Name == "svg" {
			continue
		}

		if manager.FileExists(fmt.Sprintf("svg_%s.go", tagSpec.Name)) {
			continue
		}

		source := tagGen.GenerateSource(tagSpec)

		if err := manager.WriteGoFile(fileName, source); err != nil {
			return fmt.Errorf("generate %s: %w", fileName, err)
		}
	}

	return nil
}

func generateCoreNodeFile(manager files.Manager, allSpecs []spec.TagSpec) error {
	filtered := make([]spec.TagSpec, 0, len(allSpecs))
	for _, spec := range allSpecs {
		if spec.Name == "svg" {
			continue
		}

		if manager.FileExists(fmt.Sprintf("svg_%s.go", spec.Name)) {
			continue
		}

		filtered = append(filtered, spec)
	}

	nodeGen := generator.NewNodeGenerator()
	content := nodeGen.GenerateCompleteFile(filtered)

	return manager.WriteGoFile("core_node.go", content)
}
