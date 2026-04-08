package swiftui

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"sync"

	"github.com/ebitengine/purego"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/objc"
)

// Config describes an app-owned Swift package that exposes SwiftUI content.
type Config struct {
	// Dir is the Swift package directory containing Package.swift.
	Dir string

	// Product is the dynamic library product name from Package.swift.
	Product string

	// FactoryClass is the Objective-C runtime name of the exported factory class.
	FactoryClass string

	// EmbeddedLibrary is an optional prebuilt dylib payload. When present, Open
	// materializes and loads this dylib before falling back to building from
	// source in Dir.
	EmbeddedLibrary []byte

	// EmbeddedLibraryName is the filename used when materializing
	// EmbeddedLibrary. Defaults to lib<Product>.dylib.
	EmbeddedLibraryName string

	// PrebuiltLibraryPath is an optional on-disk dylib path. When empty and Dir
	// is set, Open also checks Dir/lib<Product>.dylib before falling back to
	// swift build.
	PrebuiltLibraryPath string

	// ViewControllerSelector names the class method returning a retained
	// NSViewController. Defaults to "newRootViewController".
	ViewControllerSelector string

	// ViewSelector names the class method returning a retained NSView.
	// Defaults to "newRootView".
	ViewSelector string
}

// Bridge loads a Swift package product and resolves its SwiftUI factory class.
type Bridge struct {
	config  Config
	product string
	factory objc.Class
}

// HostedView owns a retained NSView returned by the Swift package factory.
type HostedView struct {
	appkit.NSView
	releaseOnce sync.Once
}

// HostedViewController owns a retained NSViewController returned by the Swift
// package factory.
type HostedViewController struct {
	appkit.NSViewController
	releaseOnce sync.Once
}

// Open builds and loads the Swift package product, then resolves the factory
// class in the Objective-C runtime.
func Open(cfg Config) (*Bridge, error) {
	cfg, err := normalizeConfig(cfg)
	if err != nil {
		return nil, err
	}
	var errs []error
	if len(cfg.EmbeddedLibrary) > 0 {
		path, err := materializeEmbeddedLibrary(cfg.EmbeddedLibrary, cfg.embeddedLibraryName())
		if err != nil {
			errs = append(errs, err)
		} else if bridge, err := openLoadedProduct(cfg, path); err == nil {
			return bridge, nil
		} else {
			errs = append(errs, err)
		}
	}
	if cfg.Dir != "" {
		if path, err := discoverPrebuiltDylib(cfg); err == nil {
			if bridge, err := openLoadedProduct(cfg, path); err == nil {
				return bridge, nil
			} else {
				errs = append(errs, err)
			}
		} else {
			errs = append(errs, err)
		}
		product, err := buildSwiftPackage(cfg.Dir, cfg.Product)
		if err != nil {
			errs = append(errs, err)
		} else if bridge, err := openLoadedProduct(cfg, product); err == nil {
			return bridge, nil
		} else {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		return nil, fmt.Errorf("swiftui: no embedded library or package source configured")
	}
	return nil, errors.Join(errs...)
}

// ProductPath returns the loaded dylib path.
func (b *Bridge) ProductPath() string {
	if b == nil {
		return ""
	}
	return b.product
}

// NewViewController asks the Swift factory class for a retained
// NSViewController.
func (b *Bridge) NewViewController() (*HostedViewController, error) {
	if b == nil || b.factory == 0 {
		return nil, fmt.Errorf("swiftui: nil bridge")
	}
	id := objc.Send[objc.ID](objc.ID(b.factory), objc.Sel(b.config.ViewControllerSelector))
	if id == 0 {
		return nil, fmt.Errorf("swiftui: %s.%s returned nil", b.config.FactoryClass, b.config.ViewControllerSelector)
	}
	return &HostedViewController{NSViewController: appkit.NSViewControllerFromID(id)}, nil
}

// NewView asks the Swift factory class for a retained NSView.
func (b *Bridge) NewView() (*HostedView, error) {
	if b == nil || b.factory == 0 {
		return nil, fmt.Errorf("swiftui: nil bridge")
	}
	id := objc.Send[objc.ID](objc.ID(b.factory), objc.Sel(b.config.ViewSelector))
	if id == 0 {
		return nil, fmt.Errorf("swiftui: %s.%s returned nil", b.config.FactoryClass, b.config.ViewSelector)
	}
	return &HostedView{NSView: appkit.NSViewFromID(id)}, nil
}

// Release releases the retained NSView.
func (v *HostedView) Release() {
	if v == nil {
		return
	}
	v.releaseOnce.Do(func() {
		v.NSView.Release()
		v.NSView = appkit.NSView{}
	})
}

// Release releases the retained NSViewController.
func (c *HostedViewController) Release() {
	if c == nil {
		return
	}
	c.releaseOnce.Do(func() {
		c.NSViewController.Release()
		c.NSViewController = appkit.NSViewController{}
	})
}

func normalizeConfig(cfg Config) (Config, error) {
	if runtime.GOOS != "darwin" {
		return Config{}, fmt.Errorf("swiftui: only supported on darwin")
	}
	if cfg.Product == "" {
		return Config{}, fmt.Errorf("swiftui: Product is required")
	}
	if cfg.FactoryClass == "" {
		return Config{}, fmt.Errorf("swiftui: FactoryClass is required")
	}
	if cfg.Dir != "" {
		dir, err := filepath.Abs(cfg.Dir)
		if err != nil {
			return Config{}, fmt.Errorf("swiftui: resolve dir %q: %w", cfg.Dir, err)
		}
		cfg.Dir = dir
	}
	if cfg.PrebuiltLibraryPath != "" {
		path, err := filepath.Abs(cfg.PrebuiltLibraryPath)
		if err != nil {
			return Config{}, fmt.Errorf("swiftui: resolve prebuilt dylib %q: %w", cfg.PrebuiltLibraryPath, err)
		}
		cfg.PrebuiltLibraryPath = path
	}
	if cfg.Dir == "" && len(cfg.EmbeddedLibrary) == 0 {
		return Config{}, fmt.Errorf("swiftui: Dir or EmbeddedLibrary is required")
	}
	if cfg.ViewControllerSelector == "" {
		cfg.ViewControllerSelector = "newRootViewController"
	}
	if cfg.ViewSelector == "" {
		cfg.ViewSelector = "newRootView"
	}
	return cfg, nil
}

func (cfg Config) embeddedLibraryName() string {
	if cfg.EmbeddedLibraryName != "" {
		return cfg.EmbeddedLibraryName
	}
	return "lib" + cfg.Product + ".dylib"
}

func (cfg Config) prebuiltLibraryPath() string {
	if cfg.PrebuiltLibraryPath != "" {
		return cfg.PrebuiltLibraryPath
	}
	if cfg.Dir == "" {
		return ""
	}
	return filepath.Join(cfg.Dir, cfg.embeddedLibraryName())
}

func buildSwiftPackage(dir, product string) (string, error) {
	if path, err := discoverBuiltDylib(dir, product); err == nil {
		return path, nil
	}
	if _, err := os.Stat(filepath.Join(dir, "Package.swift")); err != nil {
		return "", fmt.Errorf("swiftui: Package.swift not found in %s", dir)
	}
	cmd := exec.Command("swift", "build", "-c", "release", "--quiet", "--product", product)
	cmd.Dir = dir
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if execErr, ok := err.(*exec.Error); ok {
			return "", fmt.Errorf("swiftui: swift toolchain not found: %w", execErr)
		}
		return "", fmt.Errorf("swiftui: build %s: %w", product, err)
	}
	return discoverBuiltDylib(dir, product)
}

func openLoadedProduct(cfg Config, product string) (*Bridge, error) {
	if _, err := purego.Dlopen(product, purego.RTLD_LAZY|purego.RTLD_GLOBAL); err != nil {
		return nil, fmt.Errorf("swiftui: load %s: %w", product, err)
	}
	factory := objc.GetClass(cfg.FactoryClass)
	if factory == 0 {
		return nil, fmt.Errorf("swiftui: factory class %q not found after loading %s", cfg.FactoryClass, product)
	}
	return &Bridge{
		config:  cfg,
		product: product,
		factory: factory,
	}, nil
}

func materializeEmbeddedLibrary(data []byte, name string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("swiftui: embedded library payload is empty")
	}
	if name == "" {
		name = "libswiftui.dylib"
	}
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("swiftui: locate user cache dir: %w", err)
	}
	cacheDir = filepath.Join(cacheDir, "tmc-apple-swiftui")
	if err := os.MkdirAll(cacheDir, 0o700); err != nil {
		return "", fmt.Errorf("swiftui: create cache dir: %w", err)
	}
	sum := sha256.Sum256(data)
	path := filepath.Join(cacheDir, hex.EncodeToString(sum[:])+"-"+name)
	if existing, err := os.ReadFile(path); err == nil && sha256.Sum256(existing) == sum {
		return path, nil
	}
	tmp, err := os.CreateTemp(cacheDir, name+".*.tmp")
	if err != nil {
		return "", fmt.Errorf("swiftui: create temp dylib: %w", err)
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return "", fmt.Errorf("swiftui: write temp dylib: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return "", fmt.Errorf("swiftui: close temp dylib: %w", err)
	}
	if err := os.Chmod(tmpPath, 0o500); err != nil {
		return "", fmt.Errorf("swiftui: chmod temp dylib: %w", err)
	}
	if err := os.Rename(tmpPath, path); err != nil {
		if existing, readErr := os.ReadFile(path); readErr == nil && sha256.Sum256(existing) == sum {
			return path, nil
		}
		return "", fmt.Errorf("swiftui: finalize embedded dylib: %w", err)
	}
	return path, nil
}

func discoverPrebuiltDylib(cfg Config) (string, error) {
	path := cfg.prebuiltLibraryPath()
	if path == "" {
		return "", fmt.Errorf("swiftui: prebuilt dylib path not configured")
	}
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("swiftui: prebuilt dylib not found at %s", path)
	}
	return path, nil
}

func discoverBuiltDylib(dir, product string) (string, error) {
	name := "lib" + product + ".dylib"
	candidates := []string{
		filepath.Join(dir, ".build", "release", name),
	}
	glob, _ := filepath.Glob(filepath.Join(dir, ".build", "*", "release", name))
	sort.Strings(glob)
	candidates = append(candidates, glob...)
	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("swiftui: built product %s not found in %s", name, dir)
}
