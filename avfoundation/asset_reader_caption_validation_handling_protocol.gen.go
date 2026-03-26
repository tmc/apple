// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the methods for caption validation events.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderCaptionValidationHandling
type AVAssetReaderCaptionValidationHandling interface {
	objectivec.IObject
}

// AVAssetReaderCaptionValidationHandlingObject wraps an existing Objective-C object that conforms to the AVAssetReaderCaptionValidationHandling protocol.
type AVAssetReaderCaptionValidationHandlingObject struct {
	objectivec.Object
}
func (o AVAssetReaderCaptionValidationHandlingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAssetReaderCaptionValidationHandlingObjectFromID constructs a [AVAssetReaderCaptionValidationHandlingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAssetReaderCaptionValidationHandlingObjectFromID(id objc.ID) AVAssetReaderCaptionValidationHandlingObject {
	return AVAssetReaderCaptionValidationHandlingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the adaptor ignored one or more syntax elements
// when it created the caption object.
//
// adaptor: The adaptor object.
//
// caption: The vended caption.
//
// syntaxElements: The array of unsupported syntax elements that the adaptor object skipped.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderCaptionValidationHandling/captionAdaptor(_:didVendCaption:skippingUnsupportedSourceSyntaxElements:)
func (o AVAssetReaderCaptionValidationHandlingObject) CaptionAdaptorDidVendCaptionSkippingUnsupportedSourceSyntaxElements(adaptor objectivec.IObject, caption IAVCaption, syntaxElements []string) {
	objc.Send[struct{}](o.ID, objc.Sel("captionAdaptor:didVendCaption:skippingUnsupportedSourceSyntaxElements:"), adaptor, caption, objectivec.StringSliceToNSArray(syntaxElements))
	}

