//go:build darwin

package lldb_test

import (
	"fmt"

	"github.com/tmc/apple/x/lldb"
)

func Example() {
	d, err := lldb.Create()
	if err != nil {
		fmt.Println(err)
		return
	}
	target := d.CreateTargetWithArch("/bin/sleep", "")
	fmt.Println(target.BreakpointCreateByName("nanosleep").IsValid())
	// Output: true
}

func ExampleState_String() {
	fmt.Println(lldb.StateStopped)
	// Output: stopped
}
