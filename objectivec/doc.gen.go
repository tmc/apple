
// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

// Package objectivec provides Go bindings for the ObjectiveC framework.
//
// Gain low-level access to the Objective-C runtime and the Objective-C root types.
//
// The [Objective-C Runtime](<doc://com.apple.objectivec/documentation/ObjectiveC>) module APIs define the base of the Objective-C language. These APIs include:
//
// # Classes
//
//   - NSObject: The root class of most Objective-C class hierarchies, from which subclasses inherit a basic interface to the runtime system and the ability to behave as Objective-C objects.
//   - Protocol
//
// # Protocols
//
//   - NSObjectProtocol: The group of methods that are fundamental to all Objective-C objects.
//
// # Key Types
//
//   - [Protocol]
//
// [ObjectiveC Documentation]: https://developer.apple.com/documentation/ObjectiveC
package objectivec

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/usr/lib/libobjc.dylib"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: ObjectiveC: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

