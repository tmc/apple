// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLE5IOPort protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLE5IOPort
type MLE5IOPort interface {
	objectivec.IObject

	// BoundFeatureDirectly protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLE5IOPort/boundFeatureDirectly
	BoundFeatureDirectly() bool

	// Reset protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLE5IOPort/reset
	Reset()
}

// MLE5IOPortObject wraps an existing Objective-C object that conforms to the MLE5IOPort protocol.
type MLE5IOPortObject struct {
	objectivec.Object
}

func (o MLE5IOPortObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLE5IOPortObjectFromID constructs a [MLE5IOPortObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLE5IOPortObjectFromID(id objc.ID) MLE5IOPortObject {
	return MLE5IOPortObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5IOPort/boundFeatureDirectly
func (o MLE5IOPortObject) BoundFeatureDirectly() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("boundFeatureDirectly"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5IOPort/name
func (o MLE5IOPortObject) Name() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5IOPort/pixelBufferPool
func (o MLE5IOPortObject) PixelBufferPool() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pixelBufferPool"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5IOPort/reset
func (o MLE5IOPortObject) Reset() {
	objc.Send[struct{}](o.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5IOPort/setPixelBufferPool:
func (o MLE5IOPortObject) SetPixelBufferPool(pool objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setPixelBufferPool:"), pool)
}
