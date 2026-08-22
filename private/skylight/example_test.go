//go:build darwin

package skylight_test

import (
	"fmt"

	"github.com/tmc/apple/private/skylight"
)

func ExampleGetSLContentFilterClass() {
	cls := skylight.GetSLContentFilterClass()
	fmt.Printf("class: %T\n", cls)
	// Output:
	// class: skylight.SLContentFilterClass
}
