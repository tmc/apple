//go:build darwin

package coreservices_test

import (
	"fmt"

	"github.com/tmc/apple/coreservices"
)

func ExampleAbortErr() {
	err := coreservices.AbortErrValue
	fmt.Printf("%s (%d)\n", err, int(err))

	// Output:
	// AbortErrValue (-27)
}

func ExampleLSHandlerOptions() {
	opt := coreservices.IgnoreCreator
	fmt.Println(opt)

	// Output:
	// IgnoreCreator
}

func ExampleLSLaunchFlags() {
	flag := coreservices.AndDisplayErrors
	fmt.Println(flag)

	// Output:
	// AndDisplayErrors
}
