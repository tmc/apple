// Code generated from Apple documentation. DO NOT EDIT.

package soundanalysis

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// SNClassifierIdentifierVersion1 is version 1 of the sound classifier.
	//
	// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifierIdentifier/version1
	SNClassifierIdentifierVersion1 SNClassifierIdentifier
)

var (
	// SNErrorDomain is a string that identifies the Sound Analysis error domain.
	//
	// See: https://developer.apple.com/documentation/SoundAnalysis/SNErrorDomain
	SNErrorDomain string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SNClassifierIdentifierVersion1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SNClassifierIdentifierVersion1 = SNClassifierIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SNErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SNErrorDomain = objc.GoString(cstr)
			}
		}
	}

}
