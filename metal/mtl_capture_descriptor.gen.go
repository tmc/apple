// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLCaptureDescriptor] class.
var (
	_MTLCaptureDescriptorClass     MTLCaptureDescriptorClass
	_MTLCaptureDescriptorClassOnce sync.Once
)

func getMTLCaptureDescriptorClass() MTLCaptureDescriptorClass {
	_MTLCaptureDescriptorClassOnce.Do(func() {
		_MTLCaptureDescriptorClass = MTLCaptureDescriptorClass{class: objc.GetClass("MTLCaptureDescriptor")}
	})
	return _MTLCaptureDescriptorClass
}

// GetMTLCaptureDescriptorClass returns the class object for MTLCaptureDescriptor.
func GetMTLCaptureDescriptorClass() MTLCaptureDescriptorClass {
	return getMTLCaptureDescriptorClass()
}

type MTLCaptureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLCaptureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLCaptureDescriptorClass) Alloc() MTLCaptureDescriptor {
	rv := objc.Send[MTLCaptureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration for a Metal capture session.
//
// # Setting capture parameters
//
//   - [MTLCaptureDescriptor.CaptureObject]: The instance whose contents should be captured.
//   - [MTLCaptureDescriptor.SetCaptureObject]
//   - [MTLCaptureDescriptor.Destination]: The destination for any captured command data.
//   - [MTLCaptureDescriptor.SetDestination]
//   - [MTLCaptureDescriptor.OutputURL]: A URL for a file to write the capture data into.
//   - [MTLCaptureDescriptor.SetOutputURL]
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor
type MTLCaptureDescriptor struct {
	objectivec.Object
}

// MTLCaptureDescriptorFromID constructs a [MTLCaptureDescriptor] from an objc.ID.
//
// A configuration for a Metal capture session.
func MTLCaptureDescriptorFromID(id objc.ID) MTLCaptureDescriptor {
	return MTLCaptureDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLCaptureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLCaptureDescriptor] class.
//
// # Setting capture parameters
//
//   - [IMTLCaptureDescriptor.CaptureObject]: The instance whose contents should be captured.
//   - [IMTLCaptureDescriptor.SetCaptureObject]
//   - [IMTLCaptureDescriptor.Destination]: The destination for any captured command data.
//   - [IMTLCaptureDescriptor.SetDestination]
//   - [IMTLCaptureDescriptor.OutputURL]: A URL for a file to write the capture data into.
//   - [IMTLCaptureDescriptor.SetOutputURL]
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor
type IMTLCaptureDescriptor interface {
	objectivec.IObject

	// Topic: Setting capture parameters

	// The instance whose contents should be captured.
	CaptureObject() objectivec.IObject
	SetCaptureObject(value objectivec.IObject)
	// The destination for any captured command data.
	Destination() MTLCaptureDestination
	SetDestination(value MTLCaptureDestination)
	// A URL for a file to write the capture data into.
	OutputURL() foundation.INSURL
	SetOutputURL(value foundation.INSURL)
}

// Init initializes the instance.
func (c MTLCaptureDescriptor) Init() MTLCaptureDescriptor {
	rv := objc.Send[MTLCaptureDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLCaptureDescriptor) Autorelease() MTLCaptureDescriptor {
	rv := objc.Send[MTLCaptureDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLCaptureDescriptor creates a new MTLCaptureDescriptor instance.
func NewMTLCaptureDescriptor() MTLCaptureDescriptor {
	class := getMTLCaptureDescriptorClass()
	rv := objc.Send[MTLCaptureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The instance whose contents should be captured.
//
// # Discussion
// 
// The default value is `nil`, but you need to set an instance before using
// this descriptor to start a capture session.
// 
// The behavior of the capture session depends on the kind of instance being
// captured:
// 
// - Specify an [MTLDevice] instance to capture commands in command buffers
// created on any command queues created by the device instance. - Specify an
// [MTLCommandQueue] instance to capture commands in command buffers created
// by a specific command queue. - Specify an [MTLCaptureScope] instance to
// indirectly define which commands are captured.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/captureObject
func (c MTLCaptureDescriptor) CaptureObject() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("captureObject"))
	return objectivec.Object{ID: rv}
}
func (c MTLCaptureDescriptor) SetCaptureObject(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setCaptureObject:"), value)
}
// The destination for any captured command data.
//
// # Discussion
// 
// The default value is [CaptureDestinationDeveloperTools].
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/destination
func (c MTLCaptureDescriptor) Destination() MTLCaptureDestination {
	rv := objc.Send[MTLCaptureDestination](c.ID, objc.Sel("destination"))
	return MTLCaptureDestination(rv)
}
func (c MTLCaptureDescriptor) SetDestination(value MTLCaptureDestination) {
	objc.Send[struct{}](c.ID, objc.Sel("setDestination:"), value)
}
// A URL for a file to write the capture data into.
//
// # Discussion
// 
// The default value is `nil`. If you set [Destination] to
// [CaptureDestinationGPUTraceDocument], you need to set this property to
// where you want the file to be written to.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureDescriptor/outputURL
func (c MTLCaptureDescriptor) OutputURL() foundation.INSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (c MTLCaptureDescriptor) SetOutputURL(value foundation.INSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setOutputURL:"), value)
}

