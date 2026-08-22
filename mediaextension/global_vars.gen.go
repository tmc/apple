// Code generated from Apple documentation. DO NOT EDIT.

package mediaextension

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// MERAWProcessorReadyForMoreMediaDataDidChangeNotification is a notification that indicates a change to the object’s readiness to process additional media data.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorReadyForMoreMediaDataDidChangeNotification
	MERAWProcessorReadyForMoreMediaDataDidChangeNotification foundation.NSNotificationName
	// MERAWProcessorValuesDidChangeNotification is a notification that indicates a change to the object’s set of available processing parameters.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorValuesDidChangeNotification
	MERAWProcessorValuesDidChangeNotification foundation.NSNotificationName
	// MEVideoDecoderReadyForMoreMediaDataDidChangeNotification is a notification that indicates a change to the decoder’s readiness to process additional media data.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderReadyForMoreMediaDataDidChangeNotification
	MEVideoDecoderReadyForMoreMediaDataDidChangeNotification foundation.NSNotificationName
)

var (
	// MediaExtensionErrorDomain is the domain of the error.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MediaExtensionErrorDomain
	MediaExtensionErrorDomain foundation.NSErrorDomain
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MERAWProcessorReadyForMoreMediaDataDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MERAWProcessorReadyForMoreMediaDataDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MERAWProcessorValuesDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MERAWProcessorValuesDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MEVideoDecoderReadyForMoreMediaDataDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MEVideoDecoderReadyForMoreMediaDataDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MediaExtensionErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MediaExtensionErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

}
