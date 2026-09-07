package coremlcompiler

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteWeightRoot(t *testing.T) {
	root := t.TempDir()
	files := []WeightFile{
		{Path: "@model_path/weights/w.bin", Blob: []byte("w")},
		{Path: "@model_path/weights/b.bin", Blob: []byte("bb")},
		{Path: "@model_path/nested/dir/c.bin", Blob: []byte("ccc")},
	}
	if err := WriteWeightRoot(root, files); err != nil {
		t.Fatalf("WriteWeightRoot: %v", err)
	}
	for _, want := range []struct{ rel, data string }{
		{"weights/w.bin", "w"},
		{"weights/b.bin", "bb"},
		{"nested/dir/c.bin", "ccc"},
	} {
		got, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(want.rel)))
		if err != nil {
			t.Errorf("read %s: %v", want.rel, err)
			continue
		}
		if string(got) != want.data {
			t.Errorf("%s = %q, want %q", want.rel, got, want.data)
		}
	}
}

func TestWriteWeightRootRejects(t *testing.T) {
	for _, tt := range []struct {
		name    string
		files   []WeightFile
		wantErr string
	}{
		{
			name:    "bare relative path",
			files:   []WeightFile{{Path: "weights/w.bin"}},
			wantErr: "does not start with",
		},
		{
			name:    "prefix only",
			files:   []WeightFile{{Path: "@model_path/"}},
			wantErr: "names no file",
		},
		{
			// A generator emitting "@model_path/../x" would otherwise write
			// outside the bundle it is building.
			name:    "escapes the bundle",
			files:   []WeightFile{{Path: "@model_path/../escape.bin"}},
			wantErr: "clean relative path",
		},
		{
			name:    "unclean path",
			files:   []WeightFile{{Path: "@model_path/weights/./w.bin"}},
			wantErr: "clean relative path",
		},
		{
			// Two consts naming one file: last writer wins silently, and the
			// model loads with one of them holding the other's weights.
			name: "duplicate destination",
			files: []WeightFile{
				{Path: "@model_path/weights/w.bin", Blob: []byte("first")},
				{Path: "@model_path/weights/w.bin", Blob: []byte("second")},
			},
			wantErr: "both write",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			err := WriteWeightRoot(root, tt.files)
			if err == nil {
				t.Fatalf("WriteWeightRoot = nil error, want %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("WriteWeightRoot = %v, want %q", err, tt.wantErr)
			}
			if tt.name == "escapes the bundle" {
				if _, err := os.Stat(filepath.Join(filepath.Dir(root), "escape.bin")); err == nil {
					t.Error("wrote a file outside the weight root")
				}
			}
		})
	}
}

func TestWriteWeightRootEmptyRoot(t *testing.T) {
	if err := WriteWeightRoot("", nil); err == nil {
		t.Error("WriteWeightRoot(\"\") = nil error, want error")
	}
}

func TestWriteWeightRootRejectsSymlinkEscape(t *testing.T) {
	root, outside := t.TempDir(), t.TempDir()
	if err := os.Symlink(outside, filepath.Join(root, "weights")); err != nil {
		t.Fatal(err)
	}
	err := WriteWeightRoot(root, []WeightFile{{Path: "@model_path/weights/w.bin", Blob: []byte("weights")}})
	if err == nil {
		t.Fatal("accepted symlink outside weight root")
	}
	if _, err := os.Stat(filepath.Join(outside, "w.bin")); !os.IsNotExist(err) {
		t.Fatalf("outside file: %v", err)
	}
}
