//go:build darwin

package gtshaderprofiler_test

import (
	"fmt"

	"github.com/tmc/apple/private/xcode/gtshaderprofiler"
)

func ExampleGetGTMioTraceDataClass() {
	cls := gtshaderprofiler.GetGTMioTraceDataClass()
	fmt.Printf("class: %T\n", cls)
	// Output:
	// class: gtshaderprofiler.GTMioTraceDataClass
}
