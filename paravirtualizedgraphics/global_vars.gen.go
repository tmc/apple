// Code generated from Apple documentation. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// PGResumeErrorDomain is the error domain for suspend-resume actions.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorDomain
	PGResumeErrorDomain foundation.NSErrorDomain
)

var (
	// ParavirtualizedGraphicsVersionNumber is the framework’s current version number.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/ParavirtualizedGraphicsVersionNumber
	ParavirtualizedGraphicsVersionNumber float64
)

var (
	// ParavirtualizedGraphicsVersionString is the framework’s version number expressed as a C string.
	//
	// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/ParavirtualizedGraphicsVersionString
	ParavirtualizedGraphicsVersionString uint8
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PGResumeErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PGResumeErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ParavirtualizedGraphicsVersionNumber"); err == nil && ptr != 0 {
		ParavirtualizedGraphicsVersionNumber = objc.ValueAt[float64](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ParavirtualizedGraphicsVersionString"); err == nil && ptr != 0 {
		ParavirtualizedGraphicsVersionString = objc.ValueAt[uint8](ptr)
	}

}
