// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines a factory to create RAW processors for a codec type that the extension implements.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorExtension
type MERAWProcessorExtension interface {
	objectivec.IObject

	// A factory method to create a video RAW processor.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorExtension/makeProcessor(formatDescription:pixelBufferManager:)
	ProcessorWithFormatDescriptionExtensionPixelBufferManagerError(formatDescription coremedia.CMVideoFormatDescriptionRef, extensionPixelBufferManager IMERAWProcessorPixelBufferManager) (MERAWProcessor, error)
}

// MERAWProcessorExtensionObject wraps an existing Objective-C object that conforms to the MERAWProcessorExtension protocol.
type MERAWProcessorExtensionObject struct {
	objectivec.Object
}

func (o MERAWProcessorExtensionObject) BaseObject() objectivec.Object {
	return o.Object
}

// MERAWProcessorExtensionObjectFromID constructs a [MERAWProcessorExtensionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MERAWProcessorExtensionObjectFromID(id objc.ID) MERAWProcessorExtensionObject {
	return MERAWProcessorExtensionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A factory method to create a video RAW processor.
//
// formatDescription: A [CMVideoFormatDescription] describing the video data that was decoded to
// produce the RAW input for the video RAW processor.
//
// extensionPixelBufferManager: An [MERAWProcessorPixelBufferManager] instance that should be retained by
// the new [MERAWProcessor] instance and used for output pixelBuffer
// configuration and allocation.
//
// # Discussion
//
// Creates a new [MERAWProcessor] matching the given
// [CMVideoFormatDescriptionRef]. If these parameters are not compatible with
// the [MERAWProcessor], the call should fail, returning
// [MEErrorUnsupportedFeature].
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorExtension/makeProcessor(formatDescription:pixelBufferManager:)
func (o MERAWProcessorExtensionObject) ProcessorWithFormatDescriptionExtensionPixelBufferManagerError(formatDescription coremedia.CMVideoFormatDescriptionRef, extensionPixelBufferManager IMERAWProcessorPixelBufferManager) (MERAWProcessor, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("processorWithFormatDescription:extensionPixelBufferManager:error:"), formatDescription, extensionPixelBufferManager)
	if err != nil {
		return nil, err
	}
	return MERAWProcessorObjectFromID(rv), nil
}
