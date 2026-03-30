// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ExternalDetectedObject protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject
type ExternalDetectedObject interface {
	objectivec.IObject

	// Bounds protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/bounds
	Bounds() corefoundation.CGRect

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/center
	Center() corefoundation.CGPoint

	// ObjectType protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/objectType
	ObjectType() int64

	// SetBounds protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/setBounds:
	SetBounds(bounds corefoundation.CGRect)

	// SetCenter protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/setCenter:
	SetCenter(center corefoundation.CGPoint)

	// SetObjectType protocol.
	//
	// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/setObjectType:
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

// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/bounds
func (o ExternalDetectedObjectObject) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("bounds"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/center
func (o ExternalDetectedObjectObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/objectType
func (o ExternalDetectedObjectObject) ObjectType() int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("objectType"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/setBounds:
func (o ExternalDetectedObjectObject) SetBounds(bounds corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setBounds:"), bounds)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/setCenter:
func (o ExternalDetectedObjectObject) SetCenter(center corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), center)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/ExternalDetectedObject/setObjectType:
func (o ExternalDetectedObjectObject) SetObjectType(type_ int64) {
	objc.Send[struct{}](o.ID, objc.Sel("setObjectType:"), type_)
}
