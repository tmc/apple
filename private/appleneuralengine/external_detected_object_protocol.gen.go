// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ExternalDetectedObject protocol.
type ExternalDetectedObject interface {
	objectivec.IObject

	// Bounds protocol.
	Bounds() corefoundation.CGRect

	// Center protocol.
	Center() corefoundation.CGPoint

	// ObjectType protocol.
	ObjectType() int64

	// SetBounds protocol.
	SetBounds(bounds corefoundation.CGRect)

	// SetCenter protocol.
	SetCenter(center corefoundation.CGPoint)

	// SetObjectType protocol.
	SetObjectType(type_ int64)
}

// ExternalDetectedObjectObject wraps an existing Objective-C object that conforms to the ExternalDetectedObject protocol.
type ExternalDetectedObjectObject struct {
	objectivec.Object
}

func (o ExternalDetectedObjectObject) BaseObject() objectivec.Object {
	return o.Object
}

// ExternalDetectedObjectObjectFromID constructs a [ExternalDetectedObjectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ExternalDetectedObjectObjectFromID(id objc.ID) ExternalDetectedObjectObject {
	return ExternalDetectedObjectObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ExternalDetectedObjectObject) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("bounds"))
	return rv
}
func (o ExternalDetectedObjectObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}
func (o ExternalDetectedObjectObject) ObjectType() int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("objectType"))
	return rv
}
func (o ExternalDetectedObjectObject) SetBounds(bounds corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setBounds:"), bounds)
}
func (o ExternalDetectedObjectObject) SetCenter(center corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), center)
}
func (o ExternalDetectedObjectObject) SetObjectType(type_ int64) {
	objc.Send[struct{}](o.ID, objc.Sel("setObjectType:"), type_)
}
