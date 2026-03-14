//go:build darwin

package ane

import (
	"fmt"
	"sync"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

const (
	aneFrameworkPath     = "/System/Library/PrivateFrameworks/AppleNeuralEngine.framework/AppleNeuralEngine"
	coreMLFrameworkPath  = "/System/Library/Frameworks/CoreML.framework/CoreML"
	espressoFrameworkPath = "/System/Library/PrivateFrameworks/Espresso.framework/Espresso"
)

var (
	aneLoadOnce sync.Once
	aneLoadErr  error

	coreMLLoadOnce sync.Once
	coreMLLoadErr  error

	espressoLoadOnce sync.Once
	espressoLoadErr  error
)

// EnsureANELoaded loads the AppleNeuralEngine private framework once.
func EnsureANELoaded() error {
	aneLoadOnce.Do(func() {
		_, err := purego.Dlopen(aneFrameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err != nil {
			aneLoadErr = fmt.Errorf("load AppleNeuralEngine framework: %w", err)
		}
	})
	return aneLoadErr
}

// EnsureCoreMLLoaded loads the CoreML framework once.
func EnsureCoreMLLoaded() error {
	coreMLLoadOnce.Do(func() {
		_, err := purego.Dlopen(coreMLFrameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err != nil {
			coreMLLoadErr = fmt.Errorf("load CoreML framework: %w", err)
		}
	})
	return coreMLLoadErr
}

// EnsureEspressoLoaded loads the Espresso private framework once.
//
// Deprecated: new code should prefer [github.com/tmc/apple/x/espresso].
func EnsureEspressoLoaded() error {
	espressoLoadOnce.Do(func() {
		_, err := purego.Dlopen(espressoFrameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err != nil {
			espressoLoadErr = fmt.Errorf("load Espresso framework: %w", err)
		}
	})
	return espressoLoadErr
}

// FirstClass returns the first available Objective-C class from candidates.
func FirstClass(candidates ...string) (name string, class objc.Class) {
	for _, n := range candidates {
		if c := objc.GetClass(n); c != 0 {
			return n, c
		}
	}
	return "", 0
}
