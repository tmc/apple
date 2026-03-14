//go:build !darwin

package ane

// EnsureANELoaded loads the AppleNeuralEngine private framework once.
func EnsureANELoaded() error { return ErrNoANE }

// EnsureCoreMLLoaded loads the CoreML framework once.
func EnsureCoreMLLoaded() error { return ErrNoANE }

// EnsureEspressoLoaded loads the Espresso private framework once.
//
// Deprecated: new code should prefer [github.com/tmc/apple/x/espresso].
func EnsureEspressoLoaded() error { return ErrNoANE }
