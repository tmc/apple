//go:build darwin

package remotecoreml_test

import (
	"fmt"

	"github.com/tmc/apple/private/remotecoreml"
)

func ExampleGetCoreMLVersionClass() {
	cls := remotecoreml.GetCoreMLVersionClass()
	fmt.Printf("class: %T\n", cls)
	// Output:
	// class: remotecoreml.CoreMLVersionClass
}
