// Copyright 2026 The tmc/apple Authors. All rights reserved.

package signaturegate

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// repoRoot returns the module root: the nearest ancestor holding go.mod.
//
// It used to be filepath.Dir(wd), which encoded the assumption that this
// package sits exactly one level below the root. That is not a portable fact
// about the tree, it is a fact about where the directory happens to live, and
// moving this package one level deeper would have silently made the root
// internal/ — collapsing the population from 5821 files to zero and turning
// every gate in this package GREEN. A gate that reads the tree by path is
// exactly the thing a path change can silently narrow.
func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatalf("no go.mod in any parent of %s; the gate cannot locate the tree", dir)
		}
		dir = parent
	}
}

// generatedFiles returns every .gen.go in the tree, at any depth.
//
// It used to glob exactly <root>/<dir>/*.gen.go, one level deep, and its
// comment claimed it returned every generated file. Everything nested deeper
// was outside every gate built on it: all of private/, all of x/, and
// private/xcode/gtshaderprofiler below that. The 2026-08-18 regeneration
// dropped eight accessors from three files under private/ and no gate could
// have seen it, because those files were never in the population.
//
// A depth limit reads as a whole-tree answer, which is why this walks.
func generatedFiles(t *testing.T, root string) []string {
	t.Helper()
	var out []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		name := d.Name()
		if d.IsDir() {
			// Skip dotted directories (.git, .build) and the module cache-like
			// testdata trees; nothing generated lives in either.
			if path != root && (strings.HasPrefix(name, ".") || name == "testdata") {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(name, ".gen.go") {
			out = append(out, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(out)
	return out
}

func rel(root, path string) string {
	if r, err := filepath.Rel(root, path); err == nil {
		return r
	}
	return path
}
