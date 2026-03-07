// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An observer protocol your app implements to receive messages from the operating system’s content picker.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerObserver
type SCContentSharingPickerObserver interface {
	objectivec.IObject

	// Tells the observer that a sharing picker canceled selection for a stream.
	//
	// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerObserver/contentSharingPicker(_:didCancelFor:)
	ContentSharingPickerDidCancelForStream(picker ISCContentSharingPicker, stream ISCStream)

	// Tells the observer that a sharing picker updated the content filter for a stream.
	//
	// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerObserver/contentSharingPicker(_:didUpdateWith:for:)
	ContentSharingPickerDidUpdateWithFilterForStream(picker ISCContentSharingPicker, filter ISCContentFilter, stream ISCStream)

	// Tells the observer that a sharing picker was unable to start.
	//
	// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerObserver/contentSharingPickerStartDidFailWithError(_:)
	ContentSharingPickerStartDidFailWithError(error_ foundation.INSError)
}



// SCContentSharingPickerObserverObject wraps an existing Objective-C object that conforms to the SCContentSharingPickerObserver protocol.
type SCContentSharingPickerObserverObject struct {
	objectivec.Object
}
func (o SCContentSharingPickerObserverObject) BaseObject() objectivec.Object {
	return o.Object
}



// SCContentSharingPickerObserverObjectFromID constructs a [SCContentSharingPickerObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SCContentSharingPickerObserverObjectFromID(id objc.ID) SCContentSharingPickerObserverObject {
	return SCContentSharingPickerObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Tells the observer that a sharing picker canceled selection for a stream.
//
// picker: The content-sharing picker that sent this event.
//
// stream: The canceled stream, if any.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerObserver/contentSharingPicker(_:didCancelFor:)

func (o SCContentSharingPickerObserverObject) ContentSharingPickerDidCancelForStream(picker ISCContentSharingPicker, stream ISCStream) {
	
	objc.Send[struct{}](o.ID, objc.Sel("contentSharingPicker:didCancelForStream:"), picker, stream)
	}

// Tells the observer that a sharing picker updated the content filter for a
// stream.
//
// picker: The content-sharing picker that sent this event.
//
// filter: The new filter applied to streaming content.
//
// stream: The changed stream, if any.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerObserver/contentSharingPicker(_:didUpdateWith:for:)

func (o SCContentSharingPickerObserverObject) ContentSharingPickerDidUpdateWithFilterForStream(picker ISCContentSharingPicker, filter ISCContentFilter, stream ISCStream) {
	
	objc.Send[struct{}](o.ID, objc.Sel("contentSharingPicker:didUpdateWithFilter:forStream:"), picker, filter, stream)
	}

// Tells the observer that a sharing picker was unable to start.
//
// error: The error that caused the picker failure.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerObserver/contentSharingPickerStartDidFailWithError(_:)

func (o SCContentSharingPickerObserverObject) ContentSharingPickerStartDidFailWithError(error_ foundation.INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("contentSharingPickerStartDidFailWithError:"), error_)
	}







