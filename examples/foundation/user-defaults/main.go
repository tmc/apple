// Command user-defaults shows a small round-trip through NSUserDefaults.
//
// It writes a few values, reads them back, then removes them again.
package main

import (
	"fmt"

	"github.com/tmc/apple/foundation"
)

const (
	keyEnabled  = "example-enabled"
	keyCount    = "example-count"
	keyGreeting = "example-greeting"
)

func main() {
	defaults := foundation.GetUserDefaultsClass().StandardUserDefaults()

	// Write values.
	defaults.SetBoolForKey(true, keyEnabled)
	defaults.SetIntegerForKey(42, keyCount)

	// Store a string via NSString (implements objectivec.IObject).
	str := foundation.GetNSStringClass().Alloc().InitWithString("hello from Go")
	defaults.SetObjectForKey(str, keyGreeting)

	// Read values back.
	fmt.Printf("enabled=%t\n", defaults.BoolForKey(keyEnabled))
	fmt.Printf("count=%d\n", defaults.IntegerForKey(keyCount))
	fmt.Printf("greeting=%s\n", defaults.StringForKey(keyGreeting))

	// Clean up test keys.
	defaults.RemoveObjectForKey(keyEnabled)
	defaults.RemoveObjectForKey(keyCount)
	defaults.RemoveObjectForKey(keyGreeting)
}
