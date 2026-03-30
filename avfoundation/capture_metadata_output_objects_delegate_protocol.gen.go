// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods for receiving metadata produced by a metadata capture output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutputObjectsDelegate
type AVCaptureMetadataOutputObjectsDelegate interface {
	objectivec.IObject
}

// AVCaptureMetadataOutputObjectsDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureMetadataOutputObjectsDelegate protocol.
type AVCaptureMetadataOutputObjectsDelegateObject struct {
	objectivec.Object
}

func (o AVCaptureMetadataOutputObjectsDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureMetadataOutputObjectsDelegateObjectFromID constructs a [AVCaptureMetadataOutputObjectsDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureMetadataOutputObjectsDelegateObjectFromID(id objc.ID) AVCaptureMetadataOutputObjectsDelegateObject {
	return AVCaptureMetadataOutputObjectsDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Informs the delegate that the capture output object emitted new metadata
// objects.
//
// output: The [AVCaptureMetadataOutput] object that captured and emitted the metadata
// objects.
//
// metadataObjects: An array of [AVMetadataObject] instances representing the newly emitted
// metadata. Because [AVMetadataObject] is an abstract class, the objects in
// this array are always instances of a concrete subclass.
//
// connection: The capture connection through which the objects were emitted.
//
// # Discussion
//
// The [AVCaptureMetadataOutput] object emits only metadata objects whose
// types are included in its [MetadataObjectTypes] property. The delegate
// implements this method to perform additional processing on metadata objects
// as they become available. If you plan to use metadata objects outside the
// scope of this method, you must store strong references to them and remove
// those references when the objects are no longer required.
//
// This method is executed on the dispatch queue specified by the
// [MetadataObjectsCallbackQueue] property of the capture metadata output
// object. Because this method may be called frequently, your implementation
// should be efficient to prevent capture performance problems, including
// dropped metadata objects.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutputObjectsDelegate/metadataOutput(_:didOutput:from:)
func (o AVCaptureMetadataOutputObjectsDelegateObject) CaptureOutputDidOutputMetadataObjectsFromConnection(output IAVCaptureOutput, metadataObjects []AVMetadataObject, connection IAVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didOutputMetadataObjects:fromConnection:"), output, objectivec.IObjectSliceToNSArray(metadataObjects), connection)
}
