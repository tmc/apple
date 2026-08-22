//go:build darwin

package xcode_test

import (
	"fmt"

	"github.com/tmc/apple/private/xcode/gtshaderprofiler"
)

func Example() {
	cls := gtshaderprofiler.GetGTMioTraceDataClass()
	fmt.Printf("class: %T\n", cls)
	// Output:
	// class: gtshaderprofiler.GTMioTraceDataClass
}
