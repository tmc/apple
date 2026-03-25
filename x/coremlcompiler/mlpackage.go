package coremlcompiler

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// mlpackageManifest matches the Manifest.json format that Apple's
// MLModel.compileModelAtURL: expects.
type mlpackageManifest struct {
	FileFormatVersion   string                       `json:"fileFormatVersion"`
	ItemInfoEntries     map[string]mlpackageItemInfo  `json:"itemInfoEntries"`
	RootModelIdentifier string                       `json:"rootModelIdentifier"`
}

type mlpackageItemInfo struct {
	Author      string `json:"author"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Path        string `json:"path"`
}

// WriteMLPackage creates a .mlpackage directory at dir containing the
// given model protobuf and optional weights. The directory can be read
// back by readMLPackage or passed to MLModel.compileModelAtURL:.
//
// The layout matches Apple's coremltools format:
//
//	dir/
//	├── Manifest.json
//	└── Data/
//	    └── com.apple.CoreML/
//	        ├── model.mlmodel     (modelProto bytes)
//	        └── weights/          (copied from weightSrc if provided)
//
// Manifest paths use "com.apple.CoreML/..." (without Data/ prefix);
// the Data/ directory is implicit when resolving paths on disk.
//
// If weightSrc is non-empty and is a directory, its contents are copied
// into the package's CoreML directory preserving relative paths.
// If weightSrc is a single file, it is copied as weights/weight.bin.
func WriteMLPackage(dir string, modelProto []byte, weightSrc string) error {
	coremlDir := filepath.Join(dir, "Data", "com.apple.CoreML")
	if err := os.MkdirAll(coremlDir, 0o755); err != nil {
		return fmt.Errorf("coremlcompiler: mkdir coreml dir: %w", err)
	}

	// Generate UUIDs for manifest entries.
	modelUUID := pseudoUUID()
	manifest := mlpackageManifest{
		FileFormatVersion:   "1.0.0",
		RootModelIdentifier: modelUUID,
		ItemInfoEntries: map[string]mlpackageItemInfo{
			modelUUID: {
				Author:      "com.apple.CoreML",
				Description: "CoreML Model Specification",
				Name:        "model.mlmodel",
				Path:        "com.apple.CoreML/model.mlmodel",
			},
		},
	}

	hasWeights := false
	if weightSrc != "" {
		if info, err := os.Stat(weightSrc); err == nil {
			if info.IsDir() {
				// Check if directory is non-empty.
				entries, _ := os.ReadDir(weightSrc)
				hasWeights = len(entries) > 0
			} else {
				hasWeights = true
			}
		}
	}
	if hasWeights {
		manifest.ItemInfoEntries[pseudoUUID()] = mlpackageItemInfo{
			Author:      "com.apple.CoreML",
			Description: "CoreML Model Weights",
			Name:        "weights",
			Path:        "com.apple.CoreML/weights",
		}
	}

	manifestData, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("coremlcompiler: marshal manifest: %w", err)
	}
	manifestData = append(manifestData, '\n')
	if err := os.WriteFile(filepath.Join(dir, "Manifest.json"), manifestData, 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write manifest: %w", err)
	}

	// Write model protobuf.
	if err := os.WriteFile(filepath.Join(coremlDir, "model.mlmodel"), modelProto, 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write model.mlmodel: %w", err)
	}

	if !hasWeights {
		return nil
	}

	info, err := os.Stat(weightSrc)
	if err != nil {
		return fmt.Errorf("coremlcompiler: stat weight source: %w", err)
	}
	if info.IsDir() {
		return copyWeightDir(weightSrc, coremlDir)
	}
	// Single file -> weights/weight.bin.
	weightsDir := filepath.Join(coremlDir, "weights")
	if err := os.MkdirAll(weightsDir, 0o755); err != nil {
		return fmt.Errorf("coremlcompiler: mkdir weights: %w", err)
	}
	return copyFile(weightSrc, filepath.Join(weightsDir, "weight.bin"))
}

// pseudoUUID generates a random UUID v4 string.
func pseudoUUID() string {
	var b [16]byte
	rand.Read(b[:])
	b[6] = (b[6] & 0x0F) | 0x40 // version 4
	b[8] = (b[8] & 0x3F) | 0x80 // variant 1
	return fmt.Sprintf("%08X-%04X-%04X-%04X-%012X",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
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
