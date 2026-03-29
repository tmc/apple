// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.
// +build ios

package systemextensions

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/foundation"
)

//
// bundleID: BundleIdentifier of the application containing the system extension(s)
//
// # Return Value
// 
// A set of system extension property objects on success, nil otherwise.
//
// # Discussion
// 
// Get information about system extension(s) in an app with a bundle
// identifier
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/systemExtensions(forApplicationWithBundleID:)
func (o OSSystemExtensionsWorkspace) SystemExtensionsForApplicationWithBundleIDError(bundleID string) (foundation.INSSet, error) {
		var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("systemExtensionsForApplicationWithBundleID:error:"), objc.String(bundleID), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSSetFromID(rv), nil

}

