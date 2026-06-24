//go:build darwin

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var shimHandle uintptr

func extensionMainProbe() error {
	if err := loadSwiftShim(); err != nil {
		return err
	}
	shimClass := objc.GetClass("NinePFileSystem")
	if shimClass == 0 {
		return errors.New("swift shim did not register NinePFileSystem")
	}
	var regErr error
	objc.AutoreleasePool(func() {
		_, regErr = ensureServer(shimClass, &ninepFileSystem{config: defaultFSConfigFromEnv()})
	})
	if regErr != nil {
		return regErr
	}
	var hasClass func() bool
	if err := registerShimFunc(&hasClass, "NinePFSShimHasFileSystemClass"); err != nil {
		return err
	}
	if !hasClass() {
		return errors.New("NinePFSShimHasFileSystemClass returned false")
	}
	var runMain func()
	if err := registerShimFunc(&runMain, "NinePFSRunExtensionMain"); err != nil {
		return err
	}
	fmt.Println("9pfs: swift shim loaded")
	fmt.Println("9pfs: attached Go FSKit operation IMPs to NinePFileSystem")
	fmt.Println("9pfs: resolved NinePFSRunExtensionMain without entering ExtensionFoundation main")
	return nil
}

func registerShimFunc(fptr any, name string) error {
	sym, err := purego.Dlsym(shimHandle, name)
	if err != nil {
		return fmt.Errorf("resolve %s: %w", name, err)
	}
	purego.RegisterFunc(fptr, sym)
	return nil
}

func loadSwiftShim() error {
	path := os.Getenv("NINEPFS_SWIFTSHIM")
	if path == "" {
		var err error
		path, err = defaultSwiftShimPath()
		if err != nil {
			return err
		}
	}
	handle, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("load swift shim %s: %w", path, err)
	}
	shimHandle = handle
	return nil
}

func defaultSwiftShimPath() (string, error) {
	candidates := []string{
		filepath.Join("swiftshim", ".build", "release", "libNinePFSShim.dylib"),
		filepath.Join("swiftshim", ".build", "arm64-apple-macosx", "release", "libNinePFSShim.dylib"),
		filepath.Join("swiftshim", ".build", "x86_64-apple-macosx", "release", "libNinePFSShim.dylib"),
		filepath.Join("examples", "fskit", "9pfs", "swiftshim", ".build", "release", "libNinePFSShim.dylib"),
		filepath.Join("examples", "fskit", "9pfs", "swiftshim", ".build", "arm64-apple-macosx", "release", "libNinePFSShim.dylib"),
		filepath.Join("examples", "fskit", "9pfs", "swiftshim", ".build", "x86_64-apple-macosx", "release", "libNinePFSShim.dylib"),
	}
	if exe, err := os.Executable(); err == nil {
		contents := filepath.Dir(filepath.Dir(exe))
		candidates = append(candidates,
			filepath.Join(contents, "Frameworks", "libNinePFSShim.dylib"),
			filepath.Join(contents, "MacOS", "libNinePFSShim.dylib"),
		)
	}
	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("swift shim not built; run: (cd examples/fskit/9pfs/swiftshim && swift build -c release --product NinePFSShim)")
}
