package coremlcompiler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Compile compiles a CoreML model package (.mlpackage) or model file
// (.mlmodel) into a compiled bundle (.mlmodelc) at outputPath.
//
// Only mlprogram models (spec version 5+) are supported. Legacy
// NeuralNetwork models must be converted to the mlprogram format
// before compilation (e.g. via coremltools.convert with
// convert_to="mlprogram").
func Compile(inputPath, outputPath string) error {
	// Read the model protobuf.
	modelData, weightDir, err := readModelInput(inputPath)
	if err != nil {
		return fmt.Errorf("coremlcompiler: read input: %w", err)
	}

	// Decode the protobuf.
	model, err := decodeModel(modelData)
	if err != nil {
		return fmt.Errorf("coremlcompiler: decode: %w", err)
	}

	if model.MLProgram != nil {
		return compileMLProgram(model, weightDir, outputPath)
	}

	return fmt.Errorf("coremlcompiler: unsupported model type (only mlprogram is supported; convert NeuralNetwork models via coremltools with convert_to=\"mlprogram\")")
}

// CompileMLProgram compiles an mlprogram model from already-parsed components.
// This is useful when you have the model proto bytes and weight directory
// available separately.
func CompileMLProgram(modelProto []byte, weightDir, outputPath string) error {
	model, err := decodeModel(modelProto)
	if err != nil {
		return fmt.Errorf("coremlcompiler: decode: %w", err)
	}
	if model.MLProgram == nil {
		return fmt.Errorf("coremlcompiler: model does not contain an mlprogram")
	}
	return compileMLProgram(model, weightDir, outputPath)
}

func compileMLProgram(model *Model, weightDir, outputPath string) error {
	// Create output directory.
	if err := os.MkdirAll(outputPath, 0o755); err != nil {
		return fmt.Errorf("coremlcompiler: mkdir output: %w", err)
	}

	// 1. Emit MIL text with correct dialect for the spec version.
	milText := emitMILTextWithSpec(model.MLProgram, model.SpecVersion)
	milPath := filepath.Join(outputPath, "model.mil")
	if err := os.WriteFile(milPath, []byte(milText), 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write model.mil: %w", err)
	}

	// 2. Copy weight files.
	if weightDir != "" {
		if err := copyWeights(weightDir, outputPath, model.MLProgram); err != nil {
			return fmt.Errorf("coremlcompiler: copy weights: %w", err)
		}
	}

	// 3. Build coremldata.bin.
	coremldata := buildCoreMLData(model)
	cdPath := filepath.Join(outputPath, "coremldata.bin")
	if err := os.WriteFile(cdPath, coremldata, 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write coremldata.bin: %w", err)
	}
	if err := writeAnalyticsCoreMLData(outputPath, coremldata); err != nil {
		return fmt.Errorf("coremlcompiler: write analytics coremldata.bin: %w", err)
	}

	// 4. Build metadata.json.
	metaJSON, err := buildMetadataJSON(model)
	if err != nil {
		return fmt.Errorf("coremlcompiler: build metadata.json: %w", err)
	}
	metaPath := filepath.Join(outputPath, "metadata.json")
	if err := os.WriteFile(metaPath, metaJSON, 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write metadata.json: %w", err)
	}

	return nil
}

// CompileToTemp compiles a model to a temporary directory, returning the
// path to the compiled .mlmodelc bundle. The directory is placed under
// os.TempDir() and named by a hash of the input path for implicit caching.
func CompileToTemp(inputPath string) (string, error) {
	// Use a hash of the absolute path + modification time for cache key.
	info, err := os.Stat(inputPath)
	if err != nil {
		return "", fmt.Errorf("coremlcompiler: stat input: %w", err)
	}
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		return "", fmt.Errorf("coremlcompiler: resolve path: %w", err)
	}
	key := fmt.Sprintf("%s:%d", absPath, info.ModTime().UnixNano())
	h := fnv32a(key)
	outputPath := filepath.Join(os.TempDir(), fmt.Sprintf("coremlcompiler-%08x.mlmodelc", h))

	// Check if already compiled.
	if _, err := os.Stat(filepath.Join(outputPath, "coremldata.bin")); err == nil {
		return outputPath, nil
	}

	if err := Compile(inputPath, outputPath); err != nil {
		return "", err
	}
	return outputPath, nil
}

func fnv32a(s string) uint32 {
	var h uint32 = 2166136261
	for i := range len(s) {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

// readModelInput reads the model protobuf and locates the weight directory.
// Handles both .mlpackage (directory) and .mlmodel (single file) inputs.
func readModelInput(path string) (modelData []byte, weightDir string, err error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, "", err
	}

	if info.IsDir() {
		return readMLPackage(path)
	}

	// Single .mlmodel file — no weight directory.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	return data, filepath.Dir(path), nil
}

// readMLPackage reads an .mlpackage directory.
// The standard layout (matching coremltools output) is:
//
//	model.mlpackage/
//	├── Manifest.json
//	└── com.apple.CoreML/
//	    ├── model.mlmodel      (protobuf)
//	    └── weights/
//	        └── weight.bin     (optional)
func readMLPackage(dir string) ([]byte, string, error) {
	// Find the model file via Manifest.json.
	manifestPath := filepath.Join(dir, "Manifest.json")
	modelPath, weightDir, err := resolveManifest(manifestPath, dir)
	if err != nil {
		// Fallback: try conventional paths.
		for _, candidate := range []string{
			filepath.Join(dir, "com.apple.CoreML", "model.mlmodel"),
			filepath.Join(dir, "Data", "com.apple.CoreML", "model.mlmodel"),
		} {
			if _, serr := os.Stat(candidate); serr == nil {
				modelPath = candidate
				weightDir = filepath.Dir(candidate)
				break
			}
		}
		if modelPath == "" {
			return nil, "", fmt.Errorf("read mlpackage: no model.mlmodel found: %w", err)
		}
	}

	data, err := os.ReadFile(modelPath)
	if err != nil {
		return nil, "", fmt.Errorf("read model: %w", err)
	}

	return data, weightDir, nil
}

// resolveManifest reads Manifest.json to find the model file path.
func resolveManifest(manifestPath, packageDir string) (modelPath, weightDir string, err error) {
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return "", "", err
	}

	var manifest struct {
		RootModelIdentifier string `json:"rootModelIdentifier"`
		ItemInfoEntries     map[string]struct {
			Path   string `json:"path"`
			Name   string `json:"name"`
			Author string `json:"author"`
		} `json:"itemInfoEntries"`
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		return "", "", fmt.Errorf("parse manifest: %w", err)
	}

	// Prefer rootModelIdentifier to find the model entry.
	if manifest.RootModelIdentifier != "" {
		if entry, ok := manifest.ItemInfoEntries[manifest.RootModelIdentifier]; ok {
			modelPath = filepath.Join(packageDir, entry.Path)
			weightDir = filepath.Dir(modelPath)
			return modelPath, weightDir, nil
		}
	}

	// Fallback: find an entry whose name is "model.mlmodel".
	for _, entry := range manifest.ItemInfoEntries {
		if entry.Name == "model.mlmodel" {
			modelPath = filepath.Join(packageDir, entry.Path)
			weightDir = filepath.Dir(modelPath)
			return modelPath, weightDir, nil
		}
	}

	// Legacy fallback: find by author and append model.mlmodel.
	for _, entry := range manifest.ItemInfoEntries {
		if entry.Author == "com.apple.CoreML" {
			candidate := filepath.Join(packageDir, entry.Path)
			info, serr := os.Stat(candidate)
			if serr == nil && info.IsDir() {
				return filepath.Join(candidate, "model.mlmodel"), candidate, nil
			}
			return candidate, filepath.Dir(candidate), nil
		}
	}

	return "", "", fmt.Errorf("no CoreML entry in manifest")
}

// copyWeights copies weight files referenced by the MIL program.
// Weight references in BLOBFILE use paths like "@model_path/weights/weight.bin".
func copyWeights(srcDir, dstDir string, prog *program) error {
	// Collect all referenced blob files.
	refs := collectBlobRefs(prog)
	if len(refs) == 0 {
		return nil
	}

	for _, ref := range refs {
		// Resolve the path relative to the source directory.
		relPath := strings.TrimPrefix(ref, "@model_path/")

		srcPath := filepath.Join(srcDir, relPath)
		dstPath := filepath.Join(dstDir, relPath)

		// Skip if source doesn't exist.
		if _, err := os.Stat(srcPath); os.IsNotExist(err) {
			continue
		}

		// Create destination directory.
		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			return err
		}

		// Copy the file.
		if err := copyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("copy %s: %w", relPath, err)
		}
	}

	return nil
}

// collectBlobRefs walks the program and collects unique blob file names.
func collectBlobRefs(prog *program) []string {
	seen := make(map[string]bool)
	for _, fn := range prog.Functions {
		for _, block := range fn.BlockSpecializations {
			collectBlockBlobRefs(block, seen)
		}
	}
	refs := make([]string, 0, len(seen))
	for ref := range seen {
		refs = append(refs, ref)
	}
	return refs
}

func collectBlockBlobRefs(block *block, seen map[string]bool) {
	for _, op := range block.Operations {
		// Check operation inputs for blob references.
		for _, arg := range op.Inputs {
			for _, binding := range arg.Bindings {
				if binding.Value != nil && binding.Value.BlobFile != nil {
					seen[binding.Value.BlobFile.FileName] = true
				}
			}
		}
		// Check nested blocks.
		for _, nested := range op.Blocks {
			collectBlockBlobRefs(nested, seen)
		}
	}
}
