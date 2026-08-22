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

// validateFunctionDescriptions checks that a program with more than one MIL
// function carries the ModelDescription fields Core ML needs to address the
// extra entry points. Without them only the default function is reachable.
func validateFunctionDescriptions(model *Model) error {
	desc := model.Description
	if len(model.MLProgram.Functions) > 1 && len(desc.Functions) == 0 {
		return fmt.Errorf("multi-function program requires function descriptions")
	}
	if len(desc.Functions) > 0 && desc.DefaultFunctionName == "" {
		return fmt.Errorf("function descriptions require a default function name")
	}
	for _, fn := range desc.Functions {
		if _, ok := model.MLProgram.Functions[fn.Name]; !ok {
			return fmt.Errorf("function description %q has no matching mil function", fn.Name)
		}
	}
	if name := desc.DefaultFunctionName; name != "" {
		if _, ok := model.MLProgram.Functions[name]; !ok {
			return fmt.Errorf("default function %q has no matching mil function", name)
		}
	}
	return nil
}

func compileMLProgram(model *Model, weightDir, outputPath string) error {
	// Create output directory.
	if err := os.MkdirAll(outputPath, 0o755); err != nil {
		return fmt.Errorf("coremlcompiler: mkdir output: %w", err)
	}

	// Validate program structure.
	if err := ValidateProgram(model.MLProgram); err != nil {
		return fmt.Errorf("coremlcompiler: validate program: %w", err)
	}
	if err := validateFunctionDescriptions(model); err != nil {
		return fmt.Errorf("coremlcompiler: %w", err)
	}
	if err := validateModelProgram(model); err != nil {
		return fmt.Errorf("coremlcompiler: %w", err)
	}
	if err := validateModelInterface(model); err != nil {
		return fmt.Errorf("coremlcompiler: %w", err)
	}
	if err := validateDescriptionSignature(model); err != nil {
		return fmt.Errorf("coremlcompiler: %w", err)
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
	if err := writeAnalyticsCoreMLData(outputPath, EncodeModel(model)); err != nil {
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
// os.TempDir() and named by a hash of the input path and content for implicit caching.
func CompileToTemp(inputPath string) (string, error) {
	modelData, weightDir, err := readModelInput(inputPath)
	if err != nil {
		return "", fmt.Errorf("coremlcompiler: read input for temp compile: %w", err)
	}

	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		return "", fmt.Errorf("coremlcompiler: resolve path: %w", err)
	}

	// Compute a hash of path + model bytes + optional weight info
	var weightHash string
	if weightDir != "" {
		if entries, err := os.ReadDir(weightDir); err == nil {
			for _, e := range entries {
				if info, err := e.Info(); err == nil {
					weightHash += fmt.Sprintf(";%s:%d:%d", e.Name(), info.Size(), info.ModTime().UnixNano())
				}
			}
		}
	}
	key := fmt.Sprintf("%s;len:%d;weights:%s", absPath, len(modelData), weightHash)
	h := fnv32a(key)
	// Add digest of actual model bytes to avoid collision
	var dataHash uint32 = 2166136261
	for _, b := range modelData {
		dataHash ^= uint32(b)
		dataHash *= 16777619
	}

	outputPath := filepath.Join(os.TempDir(), fmt.Sprintf("coremlcompiler-%08x%08x.mlmodelc", h, dataHash))
	completionMarker := filepath.Join(outputPath, ".compile_complete")

	// Check if already compiled and complete.
	if _, err := os.Stat(completionMarker); err == nil {
		if _, err := os.Stat(filepath.Join(outputPath, "coremldata.bin")); err == nil {
			if _, err := os.Stat(filepath.Join(outputPath, "metadata.json")); err == nil {
				return outputPath, nil
			}
		}
	}

	// Compile into a temporary directory first.
	tmpOut, err := os.MkdirTemp(os.TempDir(), "coremlcompiler-build-*")
	if err != nil {
		return "", fmt.Errorf("coremlcompiler: create temp build dir: %w", err)
	}
	defer os.RemoveAll(tmpOut)

	model, err := decodeModel(modelData)
	if err != nil {
		return "", fmt.Errorf("coremlcompiler: decode: %w", err)
	}
	if model.MLProgram == nil {
		return "", fmt.Errorf("coremlcompiler: unsupported model type")
	}

	if err := compileMLProgram(model, weightDir, tmpOut); err != nil {
		return "", err
	}

	// Write completion marker inside temp build directory before atomic move.
	if err := os.WriteFile(filepath.Join(tmpOut, ".compile_complete"), []byte("ok"), 0o644); err != nil {
		return "", fmt.Errorf("coremlcompiler: write completion marker: %w", err)
	}

	// Atomically rename or accept existing complete directory if raced.
	if err := os.Rename(tmpOut, outputPath); err != nil {
		if _, statErr := os.Stat(completionMarker); statErr == nil {
			return outputPath, nil
		}
		// Fallback for cross-device or non-atomic file systems: directory replace if absent
		if _, statErr := os.Stat(outputPath); os.IsNotExist(statErr) {
			_ = os.Rename(tmpOut, outputPath)
		}
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

	// Manifest paths are relative (e.g. "com.apple.CoreML/model.mlmodel").
	// On disk they live under the Data/ subdirectory.
	// Helper to resolve an entry path cleanly and check it stays within packageDir.
	resolveEntry := func(entryPath string) (string, string, error) {
		// Clean and reject absolute or directory-traversing paths.
		if filepath.IsAbs(entryPath) {
			return "", "", fmt.Errorf("manifest path %q is absolute", entryPath)
		}
		rel := filepath.Clean(filepath.ToSlash(entryPath))
		if rel == ".." || strings.HasPrefix(rel, "../") || strings.Contains(rel, "/../") {
			return "", "", fmt.Errorf("manifest path %q attempts path traversal", entryPath)
		}

		cleanPkg, err := filepath.Abs(packageDir)
		if err != nil {
			return "", "", err
		}

		// Try Data/ prefix first (standard coremltools layout).
		candidate := filepath.Join(cleanPkg, "Data", filepath.FromSlash(rel))
		if _, err := os.Stat(candidate); err == nil {
			candidateAbs, err := filepath.Abs(candidate)
			if err != nil {
				return "", "", err
			}
			relToPkg, err := filepath.Rel(cleanPkg, candidateAbs)
			if err != nil || strings.HasPrefix(relToPkg, "..") {
				return "", "", fmt.Errorf("manifest path %q escapes package root", entryPath)
			}
			return candidateAbs, filepath.Dir(candidateAbs), nil
		}
		// Fall back to path as-is (non-standard layout).
		candidate = filepath.Join(cleanPkg, filepath.FromSlash(rel))
		candidateAbs, err := filepath.Abs(candidate)
		if err != nil {
			return "", "", err
		}
		relToPkg, err := filepath.Rel(cleanPkg, candidateAbs)
		if err != nil || strings.HasPrefix(relToPkg, "..") {
			return "", "", fmt.Errorf("manifest path %q escapes package root", entryPath)
		}
		return candidateAbs, filepath.Dir(candidateAbs), nil
	}

	// Prefer rootModelIdentifier to find the model entry.
	if manifest.RootModelIdentifier != "" {
		if entry, ok := manifest.ItemInfoEntries[manifest.RootModelIdentifier]; ok {
			return resolveEntry(entry.Path)
		}
	}

	// Fallback: find an entry whose name is "model.mlmodel".
	for _, entry := range manifest.ItemInfoEntries {
		if entry.Name == "model.mlmodel" {
			return resolveEntry(entry.Path)
		}
	}

	// Legacy fallback: find by author.
	for _, entry := range manifest.ItemInfoEntries {
		if entry.Author == "com.apple.CoreML" {
			return resolveEntry(entry.Path)
		}
	}

	return "", "", fmt.Errorf("no CoreML entry in manifest")
}

// copyWeights copies weight files referenced by the MIL Program.
// Weight references in BLOBFILE use paths like "@model_path/weights/weight.bin".
func copyWeights(srcDir, dstDir string, prog *Program) error {
	// Collect all referenced blob files.
	refs := collectBlobRefs(prog)
	if len(refs) == 0 {
		return nil
	}

	cleanSrc, err := filepath.Abs(srcDir)
	if err != nil {
		return fmt.Errorf("resolve weight srcDir: %w", err)
	}
	cleanDst, err := filepath.Abs(dstDir)
	if err != nil {
		return fmt.Errorf("resolve weight dstDir: %w", err)
	}

	for _, ref := range refs {
		if !strings.HasPrefix(ref, "@model_path/") {
			return fmt.Errorf("invalid weight reference %q: must start with @model_path/", ref)
		}
		// Resolve the path relative to the source directory.
		relPath := strings.TrimPrefix(ref, "@model_path/")
		if filepath.IsAbs(relPath) {
			return fmt.Errorf("invalid weight reference %q: absolute path not allowed", ref)
		}
		cleanRel := filepath.Clean(filepath.ToSlash(relPath))
		if cleanRel == ".." || strings.HasPrefix(cleanRel, "../") || strings.Contains(cleanRel, "/../") {
			return fmt.Errorf("invalid weight reference %q: path traversal forbidden", ref)
		}

		srcPath := filepath.Join(cleanSrc, filepath.FromSlash(cleanRel))
		dstPath := filepath.Join(cleanDst, filepath.FromSlash(cleanRel))

		// Ensure srcPath and dstPath do not escape cleanSrc / cleanDst
		relSrc, err := filepath.Rel(cleanSrc, srcPath)
		if err != nil || strings.HasPrefix(relSrc, "..") {
			return fmt.Errorf("weight source %q escapes source root %q", srcPath, cleanSrc)
		}
		relDst, err := filepath.Rel(cleanDst, dstPath)
		if err != nil || strings.HasPrefix(relDst, "..") {
			return fmt.Errorf("weight destination %q escapes output root %q", dstPath, cleanDst)
		}

		fi, err := os.Stat(srcPath)
		if err != nil {
			return fmt.Errorf("weight source file missing for ref %q (%s): %w", ref, srcPath, err)
		}
		if fi.IsDir() {
			return fmt.Errorf("weight source %q is a directory, expected file", srcPath)
		}

		// Create destination directory.
		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			return err
		}

		// Copy the file.
		if err := copyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("copy %s: %w", cleanRel, err)
		}
	}

	return nil
}

// collectBlobRefs walks the Program and collects unique blob file names.
func collectBlobRefs(prog *Program) []string {
	seen := make(map[string]bool)
	for _, fn := range prog.Functions {
		for _, blk := range fn.BlockSpecializations {
			collectBlockBlobRefs(blk, seen)
		}
	}
	refs := make([]string, 0, len(seen))
	for ref := range seen {
		refs = append(refs, ref)
	}
	return refs
}

func collectBlockBlobRefs(blk *Block, seen map[string]bool) {
	for _, op := range blk.Operations {
		// Check operation inputs for blob references.
		for _, arg := range op.Inputs {
			for _, b := range arg.Bindings {
				if b.Value != nil && b.Value.BlobFile != nil {
					seen[b.Value.BlobFile.FileName] = true
				}
			}
		}
		// Check operation attributes for blob references (e.g. const val).
		for _, attr := range op.Attributes {
			if attr.BlobFile != nil {
				seen[attr.BlobFile.FileName] = true
			}
		}
		// Check nested blocks.
		for _, nested := range op.Blocks {
			collectBlockBlobRefs(nested, seen)
		}
	}
}
