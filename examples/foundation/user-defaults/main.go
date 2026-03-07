package main

import (
	"fmt"

	"github.com/tmc/apple/foundation"
)

func main() {
	defaults := foundation.GetUserDefaultsClass().StandardUserDefaults()

	// Write values.
	defaults.SetBoolForKey(true, "example-enabled")
	defaults.SetIntegerForKey(42, "example-count")

	// Store a string via NSString (implements objectivec.IObject).
	str := foundation.GetNSStringClass().Alloc().InitWithString("hello from Go")
	defaults.SetObjectForKey(str, "example-greeting")

	// Read values back.
	fmt.Printf("enabled=%t\n", defaults.BoolForKey("example-enabled"))
	fmt.Printf("count=%d\n", defaults.IntegerForKey("example-count"))
	fmt.Printf("greeting=%s\n", defaults.StringForKey("example-greeting"))

	// Clean up test keys.
	defaults.RemoveObjectForKey("example-enabled")
	defaults.RemoveObjectForKey("example-count")
	defaults.RemoveObjectForKey("example-greeting")
}
