//go:build darwin

package iokit_test

import (
	"fmt"

	"github.com/tmc/apple/iokit"
)

func ExampleKIOMap() {
	fmt.Println(iokit.KIOMapAnywhere)
	fmt.Println(iokit.KIOMapReadOnly)
	fmt.Println(iokit.KIOMapStatic)

	// Output:
	// KIOMapAnywhere
	// KIOMapReadOnly
	// KIOMapStatic
}
