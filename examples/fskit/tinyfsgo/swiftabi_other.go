//go:build !darwin || !arm64

package main

import "fmt"

func tryPrivateRunningExtensionResume() error {
	return fmt.Errorf("private running-extension resume probe requires darwin/arm64")
}
