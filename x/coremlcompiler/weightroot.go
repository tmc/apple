package coremlcompiler

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// ModelPathPrefix is the prefix a BLOBFILE path in MIL text uses to name a
// location inside the compiled bundle. The runtime resolves it against the
// bundle directory, so "@model_path/weights/w.bin" is the file "weights/w.bin"
// in the .mlmodelc.
const ModelPathPrefix = "@model_path/"

// A WeightFile is one weight payload addressed by the MIL text that reads it.
// Path is the BLOBFILE path as it appears in the MIL text, normally rooted at
// [ModelPathPrefix]; Blob is the file's contents.
type WeightFile struct {
	Path string
	Blob []byte
}

// WriteWeightRoot writes files into root under the relative paths their MIL
// BLOBFILE paths name, creating directories as needed.
//
// root is then a weight root suitable for [CompileMILText], which copies it
// into the compiled bundle so that "@model_path/..." resolves at load time.
func WriteWeightRoot(root string, files []WeightFile) error {
	if root == "" {
		return fmt.Errorf("coremlcompiler: weight root is empty")
	}
	if err := os.MkdirAll(root, 0o755); err != nil {
		return fmt.Errorf("coremlcompiler: create weight root: %w", err)
	}
	dir, err := os.OpenRoot(root)
	if err != nil {
		return fmt.Errorf("coremlcompiler: open weight root: %w", err)
	}
	defer dir.Close()
	seen := make(map[string]int, len(files))
	for i, f := range files {
		rel, err := relativeWeightPath(f.Path)
		if err != nil {
			return fmt.Errorf("coremlcompiler: weight file %q: %w", f.Path, err)
		}
		// Two consts writing one path would leave whichever blob landed last,
		// which reads at runtime as the wrong weights rather than as an error.
		if j, dup := seen[rel]; dup {
			return fmt.Errorf("coremlcompiler: weight files %d and %d both write %q", j, i, rel)
		}
		seen[rel] = i

		dst := filepath.FromSlash(rel)
		if err := dir.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return fmt.Errorf("coremlcompiler: mkdir for weight %q: %w", f.Path, err)
		}
		if err := dir.WriteFile(dst, f.Blob, 0o644); err != nil {
			return fmt.Errorf("coremlcompiler: write weight %q: %w", f.Path, err)
		}
	}
	return nil
}

// relativeWeightPath converts a MIL BLOBFILE path to a slash-separated path
// relative to the bundle root.
func relativeWeightPath(p string) (string, error) {
	if !strings.HasPrefix(p, ModelPathPrefix) {
		return "", fmt.Errorf("path does not start with %q", ModelPathPrefix)
	}
	rel := strings.TrimPrefix(p, ModelPathPrefix)
	if rel == "" {
		return "", fmt.Errorf("path names no file")
	}
	// The path comes from generated MIL text, so it is not attacker-controlled,
	// but it does decide where this function writes; keep it inside the root.
	clean := path.Clean(rel)
	if clean != rel || path.IsAbs(clean) || clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return "", fmt.Errorf("path is not a clean relative path inside the bundle")
	}
	return clean, nil
}
