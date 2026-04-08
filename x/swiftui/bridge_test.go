package swiftui

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNormalizeConfigDefaults(t *testing.T) {
	dir := t.TempDir()
	cfg, err := normalizeConfig(Config{
		Dir:          dir,
		Product:      "ExampleUI",
		FactoryClass: "ExampleUIBridge",
	})
	if err != nil {
		t.Fatalf("normalizeConfig() error = %v", err)
	}
	if cfg.ViewControllerSelector != "newRootViewController" {
		t.Fatalf("ViewControllerSelector = %q, want %q", cfg.ViewControllerSelector, "newRootViewController")
	}
	if cfg.ViewSelector != "newRootView" {
		t.Fatalf("ViewSelector = %q, want %q", cfg.ViewSelector, "newRootView")
	}
	if !filepath.IsAbs(cfg.Dir) {
		t.Fatalf("Dir = %q, want absolute path", cfg.Dir)
	}
}

func TestDiscoverBuiltDylib(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".build", "release", "libExampleUI.dylib")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("test"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := discoverBuiltDylib(dir, "ExampleUI")
	if err != nil {
		t.Fatalf("discoverBuiltDylib() error = %v", err)
	}
	if got != path {
		t.Fatalf("discoverBuiltDylib() = %q, want %q", got, path)
	}
}

func TestDiscoverPrebuiltDylibDefaultPath(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "libExampleUI.dylib")
	if err := os.WriteFile(path, []byte("test"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := discoverPrebuiltDylib(Config{
		Dir:     dir,
		Product: "ExampleUI",
	})
	if err != nil {
		t.Fatalf("discoverPrebuiltDylib() error = %v", err)
	}
	if got != path {
		t.Fatalf("discoverPrebuiltDylib() = %q, want %q", got, path)
	}
}

func TestMaterializeEmbeddedLibrary(t *testing.T) {
	path, err := materializeEmbeddedLibrary([]byte("test-dylib"), "libExampleUI.dylib")
	if err != nil {
		t.Fatalf("materializeEmbeddedLibrary() error = %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q) error = %v", path, err)
	}
	if string(data) != "test-dylib" {
		t.Fatalf("materialized data = %q, want %q", string(data), "test-dylib")
	}
	path2, err := materializeEmbeddedLibrary([]byte("test-dylib"), "libExampleUI.dylib")
	if err != nil {
		t.Fatalf("second materializeEmbeddedLibrary() error = %v", err)
	}
	if path2 != path {
		t.Fatalf("materializeEmbeddedLibrary() path = %q, want %q", path2, path)
	}
}
