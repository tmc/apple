//go:build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	const (
		product = "ExampleAppUI"
		output  = "libExampleAppUI.dylib"
	)

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fail("resolve generator path")
	}
	root := filepath.Dir(file)

	armScratch := filepath.Join(root, ".build", "universal-arm64")
	amdScratch := filepath.Join(root, ".build", "universal-x86_64")
	armLib := filepath.Join(armScratch, "arm64-apple-macosx", "release", output)
	amdLib := filepath.Join(amdScratch, "x86_64-apple-macosx", "release", output)
	universalLib := filepath.Join(root, output)

	_ = os.RemoveAll(armScratch)
	_ = os.RemoveAll(amdScratch)

	fmt.Println("building arm64 release dylib...")
	run(root, "swift", "build", "-c", "release", "--quiet", "--product", product,
		"--triple", "arm64-apple-macosx",
		"--scratch-path", armScratch,
	)

	fmt.Println("building x86_64 release dylib...")
	run(root, "swift", "build", "-c", "release", "--quiet", "--product", product,
		"--triple", "x86_64-apple-macosx",
		"--scratch-path", amdScratch,
	)

	fmt.Println("creating universal dylib...")
	run(root, "lipo", "-create", "-output", universalLib, armLib, amdLib)

	fmt.Printf("wrote %s\n", universalLib)
}

func run(dir, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fail("%s %v: %v", name, args, err)
	}
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
