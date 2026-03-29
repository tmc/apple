// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureMetadataOutput] class.
var (
	_AVCaptureMetadataOutputClass     AVCaptureMetadataOutputClass
	_AVCaptureMetadataOutputClassOnce sync.Once
)

func getAVCaptureMetadataOutputClass() AVCaptureMetadataOutputClass {
	_AVCaptureMetadataOutputClassOnce.Do(func() {
		_AVCaptureMetadataOutputClass = AVCaptureMetadataOutputClass{class: objc.GetClass("AVCaptureMetadataOutput")}
	})
	return _AVCaptureMetadataOutputClass
}

// GetAVCaptureMetadataOutputClass returns the class object for AVCaptureMetadataOutput.
func GetAVCaptureMetadataOutputClass() AVCaptureMetadataOutputClass {
	return getAVCaptureMetadataOutputClass()
}

type AVCaptureMetadataOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureMetadataOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureMetadataOutputClass) Alloc() AVCaptureMetadataOutput {
	rv := objc.Send[AVCaptureMetadataOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A capture output for processing timed metadata produced by a capture
// session.
//
// # Overview
// 
// An [AVCaptureMetadataOutput] object intercepts metadata objects emitted by
// its associated capture connection and forwards them to a delegate object
// for processing. You can use instances of this class to process specific
// types of metadata included with the input data. You use this class the way
// you do other output objects, typically by adding it as an output to an
// [AVCaptureSession] object.
//
// # Configuring metadata capture
//
//   - [AVCaptureMetadataOutput.AvailableMetadataObjectTypes]: An array of strings identifying the types of metadata objects that can be captured.
//   - [AVCaptureMetadataOutput.MetadataObjectTypes]: An array of strings identifying the types of metadata objects  to process.
//   - [AVCaptureMetadataOutput.SetMetadataObjectTypes]
//   - [AVCaptureMetadataOutput.RectOfInterest]: A rectangle of interest for limiting the search area for visual metadata.
//   - [AVCaptureMetadataOutput.SetRectOfInterest]
//   - [AVCaptureMetadataOutput.RequiredMetadataObjectTypesForCinematicVideoCapture]: The required metadata object types when Cinematic Video capture is enabled.
//
// # Receiving captured metadata objects
//
//   - [AVCaptureMetadataOutput.SetMetadataObjectsDelegateQueue]: Sets the delegate and dispatch queue to use handle callbacks.
//   - [AVCaptureMetadataOutput.MetadataObjectsDelegate]: The delegate of the capture metadata output object.
//   - [AVCaptureMetadataOutput.MetadataObjectsCallbackQueue]: The dispatch queue on which to execute the delegate’s methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput
type AVCaptureMetadataOutput struct {
	AVCaptureOutput
}

// AVCaptureMetadataOutputFromID constructs a [AVCaptureMetadataOutput] from an objc.ID.
//
// A capture output for processing timed metadata produced by a capture
// session.
func AVCaptureMetadataOutputFromID(id objc.ID) AVCaptureMetadataOutput {
	return AVCaptureMetadataOutput{AVCaptureOutput: AVCaptureOutputFromID(id)}
}
// NOTE: AVCaptureMetadataOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureMetadataOutput] class.
//
// # Configuring metadata capture
//
//   - [IAVCaptureMetadataOutput.AvailableMetadataObjectTypes]: An array of strings identifying the types of metadata objects that can be captured.
//   - [IAVCaptureMetadataOutput.MetadataObjectTypes]: An array of strings identifying the types of metadata objects  to process.
//   - [IAVCaptureMetadataOutput.SetMetadataObjectTypes]
//   - [IAVCaptureMetadataOutput.RectOfInterest]: A rectangle of interest for limiting the search area for visual metadata.
//   - [IAVCaptureMetadataOutput.SetRectOfInterest]
//   - [IAVCaptureMetadataOutput.RequiredMetadataObjectTypesForCinematicVideoCapture]: The required metadata object types when Cinematic Video capture is enabled.
//
// # Receiving captured metadata objects
//
//   - [IAVCaptureMetadataOutput.SetMetadataObjectsDelegateQueue]: Sets the delegate and dispatch queue to use handle callbacks.
//   - [IAVCaptureMetadataOutput.MetadataObjectsDelegate]: The delegate of the capture metadata output object.
//   - [IAVCaptureMetadataOutput.MetadataObjectsCallbackQueue]: The dispatch queue on which to execute the delegate’s methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput
type IAVCaptureMetadataOutput interface {
	IAVCaptureOutput

	// Topic: Configuring metadata capture

	// An array of strings identifying the types of metadata objects that can be captured.
	AvailableMetadataObjectTypes() []string
	// An array of strings identifying the types of metadata objects  to process.
	MetadataObjectTypes() []string
	SetMetadataObjectTypes(value []string)
	// A rectangle of interest for limiting the search area for visual metadata.
	RectOfInterest() corefoundation.CGRect
	SetRectOfInterest(value corefoundation.CGRect)
	// The required metadata object types when Cinematic Video capture is enabled.
	RequiredMetadataObjectTypesForCinematicVideoCapture() []string

	// Topic: Receiving captured metadata objects

	// Sets the delegate and dispatch queue to use handle callbacks.
	SetMetadataObjectsDelegateQueue(objectsDelegate AVCaptureMetadataOutputObjectsDelegate, objectsCallbackQueue dispatch.Queue)
	// The delegate of the capture metadata output object.
	MetadataObjectsDelegate() AVCaptureMetadataOutputObjectsDelegate
	// The dispatch queue on which to execute the delegate’s methods.
	MetadataObjectsCallbackQueue() dispatch.Queue
}

// Init initializes the instance.
func (c AVCaptureMetadataOutput) Init() AVCaptureMetadataOutput {
	rv := objc.Send[AVCaptureMetadataOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureMetadataOutput) Autorelease() AVCaptureMetadataOutput {
	rv := objc.Send[AVCaptureMetadataOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureMetadataOutput creates a new AVCaptureMetadataOutput instance.
func NewAVCaptureMetadataOutput() AVCaptureMetadataOutput {
	class := getAVCaptureMetadataOutputClass()
	rv := objc.Send[AVCaptureMetadataOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the delegate and dispatch queue to use handle callbacks.
//
// objectsDelegate: The delegate object to notify when new metadata objects become available.
// This object must conform to the [AVCaptureMetadataOutputObjectsDelegate]
// protocol.
//
// objectsCallbackQueue: The dispatch queue on which to execute the delegate’s methods. This queue
// must be a serial queue to ensure that metadata objects are delivered in the
// order in which they were received. If the `objectsDelegate` parameter is
// `nil`, you may specify `nil` for this parameter too; otherwise, you must
// specify a valid dispatch queue.
//
// # Discussion
// 
// When new metadata objects are captured from the receiver’s connection,
// they are vended to the delegate object. All delegate methods are executed
// on the dispatch queue specified in the `objectsCallbackQueue` parameter. To
// ensure that metadata objects are processed in a timely manner and not
// dropped, you should specify a dispatch queue dedicated to processing the
// objects or that is otherwise not busy.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/setMetadataObjectsDelegate(_:queue:)
func (c AVCaptureMetadataOutput) SetMetadataObjectsDelegateQueue(objectsDelegate AVCaptureMetadataOutputObjectsDelegate, objectsCallbackQueue dispatch.Queue) {
	objc.Send[objc.ID](c.ID, objc.Sel("setMetadataObjectsDelegate:queue:"), objectsDelegate, uintptr(objectsCallbackQueue.Handle()))
}

// An array of strings identifying the types of metadata objects that can be
// captured.
//
// # Discussion
// 
// Each string in the array corresponds to a possible value in the [Type]
// property of the [AVMetadataObject] objects reported by the receiver. The
// available types are dependent on the capabilities of the
// [AVCaptureInputPort] to which the receiver’s connection is attached.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/availableMetadataObjectTypes
func (c AVCaptureMetadataOutput) AvailableMetadataObjectTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableMetadataObjectTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// An array of strings identifying the types of metadata objects to process.
//
// # Discussion
// 
// This property is used to filter the metadata objects reported by the
// receiver. Only metadata objects whose type matches one of the strings in
// this property are forwarded to the delegate’s
// [CaptureOutputDidOutputMetadataObjectsFromConnection] method for
// processing.
// 
// When assigning a new array to this property, each of the type strings must
// be present in the array returned by the [AvailableMetadataObjectTypes]
// property; otherwise, the receiver raises an[NSException].
// 
// The default is an empty [NSArray] object, and as a result, no metadata
// objects are forwarded to the delegate’s
// [CaptureOutputDidOutputMetadataObjectsFromConnection] method. The same
// result can be achieved by setting the property to `nil`. This default
// behavior maximizes both performance and battery life.
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSException]: https://developer.apple.com/documentation/Foundation/NSException
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectTypes
func (c AVCaptureMetadataOutput) MetadataObjectTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("metadataObjectTypes"))
	return objc.ConvertSliceToStrings(rv)
}
func (c AVCaptureMetadataOutput) SetMetadataObjectTypes(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setMetadataObjectTypes:"), objectivec.StringSliceToNSArray(value))
}
// A rectangle of interest for limiting the search area for visual metadata.
//
// # Discussion
// 
// The value of this property is a [CGRect] value that determines the
// object’s rectangle of interest for each frame of video.
// 
// The rectangle’s origin is top left and is relative to the coordinate
// space of the device providing the metadata.
// 
// Specifying a rectangle of interest may improve detection performance for
// certain types of metadata. Metadata objects whose bounds do not intersect
// with the `rectOfInterest` will not be returned.
// 
// The default value of this property is a rectangle of `(0.0, 0.0, 1.0,
// 1.0)`.
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/rectOfInterest
func (c AVCaptureMetadataOutput) RectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("rectOfInterest"))
	return corefoundation.CGRect(rv)
}
func (c AVCaptureMetadataOutput) SetRectOfInterest(value corefoundation.CGRect) {
	objc.Send[struct{}](c.ID, objc.Sel("setRectOfInterest:"), value)
}
// The required metadata object types when Cinematic Video capture is enabled.
//
// # Discussion
// 
// Since the Cinematic Video algorithm requires a particular set of metadata
// objects to function optimally, you must set your [MetadataObjectTypes]
// property to this property’s returned value if you’ve set
// [CinematicVideoCaptureEnabled] to `true` on the connected device input,
// otherwise an [NSInvalidArgumentException] is thrown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/requiredMetadataObjectTypesForCinematicVideoCapture
func (c AVCaptureMetadataOutput) RequiredMetadataObjectTypesForCinematicVideoCapture() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("requiredMetadataObjectTypesForCinematicVideoCapture"))
	return objc.ConvertSliceToStrings(rv)
}
// The delegate of the capture metadata output object.
//
// # Discussion
// 
// The delegate object must conform to the
// [AVCaptureMetadataOutputObjectsDelegate] protocol. The object in this
// property is used to process all metadata objects captured from the capture
// metadata output object’s connection.
// 
// To set the delegate object, you must use the
// [SetMetadataObjectsDelegateQueue] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectsDelegate
func (c AVCaptureMetadataOutput) MetadataObjectsDelegate() AVCaptureMetadataOutputObjectsDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("metadataObjectsDelegate"))
	return AVCaptureMetadataOutputObjectsDelegateObjectFromID(rv)
}
// The dispatch queue on which to execute the delegate’s methods.
//
// # Discussion
// 
// To set the dispatch queue, you must use the
// [SetMetadataObjectsDelegateQueue] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput/metadataObjectsCallbackQueue
func (c AVCaptureMetadataOutput) MetadataObjectsCallbackQueue() dispatch.Queue {
	rv := objc.Send[uintptr](c.ID, objc.Sel("metadataObjectsCallbackQueue"))
	return dispatch.QueueFromHandle(rv)
}

