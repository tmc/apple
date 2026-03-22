// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A specialized memory buffer that stores a GPU’s counter set data.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer
type MTLCounterSampleBuffer interface {
	objectivec.IObject

	// A string that identifies the counter sample buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer/label
	Label() string

	// The GPU device instance that owns the counter sample buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer/device
	Device() MTLDevice

	// The number of samples in the buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer/sampleCount
	SampleCount() uint

	// Transforms samples of a GPU’s counter set from the driver’s internal format to a standard Metal data structure.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer/resolveCounterRange:
	ResolveCounterRange(range_ foundation.NSRange) foundation.INSData
}

// MTLCounterSampleBufferObject wraps an existing Objective-C object that conforms to the MTLCounterSampleBuffer protocol.
type MTLCounterSampleBufferObject struct {
	objectivec.Object
}
func (o MTLCounterSampleBufferObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLCounterSampleBufferObjectFromID constructs a [MTLCounterSampleBufferObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLCounterSampleBufferObjectFromID(id objc.ID) MTLCounterSampleBufferObject {
	return MTLCounterSampleBufferObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A string that identifies the counter sample buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer/label
func (o MTLCounterSampleBufferObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// The GPU device instance that owns the counter sample buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer/device
func (o MTLCounterSampleBufferObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// The number of samples in the buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer/sampleCount
func (o MTLCounterSampleBufferObject) SampleCount() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("sampleCount"))
	return rv
	}
// Transforms samples of a GPU’s counter set from the driver’s internal
// format to a standard Metal data structure.
//
// range: A range that indicates which sample instances the method resolves in the
// counter sample buffer.
//
// # Return Value
// 
// An [NSData] instance if the method successfully resolves the range of
// samples in the buffer; otherwise, `nil`.
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// # Discussion
// 
// You can only call this method on a counter sample buffer that you create
// with [StorageModeShared] (see [StorageMode]). For an example of how and
// when to use this method, see [Converting a GPU’s counter data into a
// readable format].
//
// [Converting a GPU’s counter data into a readable format]: https://developer.apple.com/documentation/Metal/converting-a-gpus-counter-data-into-a-readable-format
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBuffer/resolveCounterRange:
func (o MTLCounterSampleBufferObject) ResolveCounterRange(range_ foundation.NSRange) foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("resolveCounterRange:"), range_)
	return foundation.NSDataFromID(rv)
	}

