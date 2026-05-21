// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLE5PortBinder protocol.
type MLE5PortBinder interface {
	objectivec.IObject

	// Reset protocol.
	Reset()
}

// MLE5PortBinderObject wraps an existing Objective-C object that conforms to the MLE5PortBinder protocol.
type MLE5PortBinderObject struct {
	objectivec.Object
}

func (o MLE5PortBinderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLE5PortBinderObjectFromID constructs a [MLE5PortBinderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLE5PortBinderObjectFromID(id objc.ID) MLE5PortBinderObject {
	return MLE5PortBinderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLE5PortBinderObject) PixelBufferPool() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pixelBufferPool"))
	return objectivec.Object{ID: rv}
}
func (o MLE5PortBinderObject) Reset() {
	objc.Send[struct{}](o.ID, objc.Sel("reset"))
}
func (o MLE5PortBinderObject) SetPixelBufferPool(pool objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setPixelBufferPool:"), pool)
}
