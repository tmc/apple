// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLE5IOPort protocol.
type MLE5IOPort interface {
	objectivec.IObject

	// BoundFeatureDirectly protocol.
	BoundFeatureDirectly() bool

	// Name protocol.
	Name() objectivec.IObject

	// PixelBufferPool protocol.
	PixelBufferPool() objectivec.IObject

	// Reset protocol.
	Reset()

	// SetPixelBufferPool protocol.
	SetPixelBufferPool(pool objectivec.IObject)
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

func (o MLE5IOPortObject) BoundFeatureDirectly() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("boundFeatureDirectly"))
	return rv
}
func (o MLE5IOPortObject) Name() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("name"))
	return objectivec.Object{ID: rv}
}
func (o MLE5IOPortObject) PixelBufferPool() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("pixelBufferPool"))
	return objectivec.Object{ID: rv}
}
func (o MLE5IOPortObject) Reset() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("reset"))
}
func (o MLE5IOPortObject) SetPixelBufferPool(pool objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setPixelBufferPool:"), pool)
}
