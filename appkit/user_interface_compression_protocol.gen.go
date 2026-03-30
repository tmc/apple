// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that describes how a UI control should redisplay when space is restricted.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompression
type NSUserInterfaceCompression interface {
	objectivec.IObject

	// Compress the view by applying the specified compression options.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompression/compress(withPrioritizedCompressionOptions:)
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []NSUserInterfaceCompressionOptions)

	// Returns the minimum size a view can achieve by applying the supplied compression options.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompression/minimumSize(withPrioritizedCompressionOptions:)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []NSUserInterfaceCompressionOptions) corefoundation.CGSize

	// The compression options that are currently applied to the view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompression/activeCompressionOptions
	ActiveCompressionOptions() INSUserInterfaceCompressionOptions
}

// NSUserInterfaceCompressionObject wraps an existing Objective-C object that conforms to the NSUserInterfaceCompression protocol.
type NSUserInterfaceCompressionObject struct {
	objectivec.Object
}

func (o NSUserInterfaceCompressionObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSUserInterfaceCompressionObjectFromID constructs a [NSUserInterfaceCompressionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSUserInterfaceCompressionObjectFromID(id objc.ID) NSUserInterfaceCompressionObject {
	return NSUserInterfaceCompressionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Compress the view by applying the specified compression options.
//
// prioritizedOptions: An array of compression options that the view should apply to reduce its
// size.
//
// # Discussion
//
// When the system calls this method, the view should resize itself according
// to the compression options supplied.
//
// Compression options that are handled by the system are not included in the
// supplied array.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompression/compress(withPrioritizedCompressionOptions:)
func (o NSUserInterfaceCompressionObject) CompressWithPrioritizedCompressionOptions(prioritizedOptions []NSUserInterfaceCompressionOptions) {
	objc.Send[struct{}](o.ID, objc.Sel("compressWithPrioritizedCompressionOptions:"), objectivec.IObjectSliceToNSArray(prioritizedOptions))
}

// Returns the minimum size a view can achieve by applying the supplied
// compression options.
//
// prioritizedOptions: An array of compression options that the view should apply to reduce its
// size.
//
// # Return Value
//
// The minimum size of a view when applying the supplied compression options.
//
// # Discussion
//
// Compression options that are handled by the system are not included in the
// supplied array.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompression/minimumSize(withPrioritizedCompressionOptions:)
func (o NSUserInterfaceCompressionObject) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []NSUserInterfaceCompressionOptions) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), objectivec.IObjectSliceToNSArray(prioritizedOptions))
	return rv
}

// The compression options that are currently applied to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompression/activeCompressionOptions
func (o NSUserInterfaceCompressionObject) ActiveCompressionOptions() INSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("activeCompressionOptions"))
	return NSUserInterfaceCompressionOptionsFromID(rv)
}
