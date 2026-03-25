package coremlcompiler

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// WriteMLPackage creates a .mlpackage directory at dir containing the
// given model protobuf and optional weights. The directory can be read
// back by readMLPackage or passed to MLModel.compileModelAtURL:.
//
// If weightSrc is non-empty and is a directory, its contents are copied
// into the package's data directory preserving relative paths (so a file
// at weightSrc/weights/weight.bin ends up at Data/com.apple.CoreML/weights/weight.bin).
// If weightSrc is a single file, it is copied as Data/com.apple.CoreML/weights/weight.bin.
func WriteMLPackage(dir string, modelProto []byte, weightSrc string) error {
	dataDir := filepath.Join(dir, "Data", "com.apple.CoreML")
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return fmt.Errorf("coremlcompiler: mkdir data: %w", err)
	}

	// Write Manifest.json.
	manifest := struct {
		FileFormatVersion string `json:"fileFormatVersion"`
		ItemInfoEntries   map[string]struct {
			Path   string `json:"path"`
			Author string `json:"author"`
		} `json:"itemInfoEntries"`
	}{
		FileFormatVersion: "1.0.0",
		ItemInfoEntries: map[string]struct {
			Path   string `json:"path"`
			Author string `json:"author"`
		}{
			"com.apple.CoreML": {
				Path:   "Data/com.apple.CoreML",
				Author: "com.apple.CoreML",
			},
		},
	}
	manifestData, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("coremlcompiler: marshal manifest: %w", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "Manifest.json"), manifestData, 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write manifest: %w", err)
	}

	// Write model protobuf.
	if err := os.WriteFile(filepath.Join(dataDir, "model.mlmodel"), modelProto, 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write model.mlmodel: %w", err)
	}

	// Copy weights if provided.
	if weightSrc == "" {
		return nil
	}
	info, err := os.Stat(weightSrc)
	if err != nil {
		return fmt.Errorf("coremlcompiler: stat weight source: %w", err)
	}
	if info.IsDir() {
		return copyWeightDir(weightSrc, dataDir)
	}
	// Single file -> weights/weight.bin.
	weightsDir := filepath.Join(dataDir, "weights")
	if err := os.MkdirAll(weightsDir, 0o755); err != nil {
		return fmt.Errorf("coremlcompiler: mkdir weights: %w", err)
	}
	return copyFile(weightSrc, filepath.Join(weightsDir, "weight.bin"))
}

// copyWeightDir copies all files from src into dst preserving relative paths.
func copyWeightDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		return copyFile(path, target)
	})
}
