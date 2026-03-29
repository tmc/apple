// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNRecognizedPoint3D] class.
var (
	_VNRecognizedPoint3DClass     VNRecognizedPoint3DClass
	_VNRecognizedPoint3DClassOnce sync.Once
)

func getVNRecognizedPoint3DClass() VNRecognizedPoint3DClass {
	_VNRecognizedPoint3DClassOnce.Do(func() {
		_VNRecognizedPoint3DClass = VNRecognizedPoint3DClass{class: objc.GetClass("VNRecognizedPoint3D")}
	})
	return _VNRecognizedPoint3DClass
}

// GetVNRecognizedPoint3DClass returns the class object for VNRecognizedPoint3D.
func GetVNRecognizedPoint3DClass() VNRecognizedPoint3DClass {
	return getVNRecognizedPoint3DClass()
}

type VNRecognizedPoint3DClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNRecognizedPoint3DClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizedPoint3DClass) Alloc() VNRecognizedPoint3D {
	rv := objc.Send[VNRecognizedPoint3D](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A 3D point that includes an identifier to the point.
//
// # Getting the Identifier
//
//   - [VNRecognizedPoint3D.Identifier]: The identifier that provides context about what kind of point the request recognizes.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoint3D
type VNRecognizedPoint3D struct {
	VNPoint3D
}

// VNRecognizedPoint3DFromID constructs a [VNRecognizedPoint3D] from an objc.ID.
//
// A 3D point that includes an identifier to the point.
func VNRecognizedPoint3DFromID(id objc.ID) VNRecognizedPoint3D {
	return VNRecognizedPoint3D{VNPoint3D: VNPoint3DFromID(id)}
}
// NOTE: VNRecognizedPoint3D adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNRecognizedPoint3D] class.
//
// # Getting the Identifier
//
//   - [IVNRecognizedPoint3D.Identifier]: The identifier that provides context about what kind of point the request recognizes.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoint3D
type IVNRecognizedPoint3D interface {
	IVNPoint3D

	// Topic: Getting the Identifier

	// The identifier that provides context about what kind of point the request recognizes.
	Identifier() VNRecognizedPointKey
}

// Init initializes the instance.
func (r VNRecognizedPoint3D) Init() VNRecognizedPoint3D {
	rv := objc.Send[VNRecognizedPoint3D](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizedPoint3D) Autorelease() VNRecognizedPoint3D {
	rv := objc.Send[VNRecognizedPoint3D](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizedPoint3D creates a new VNRecognizedPoint3D instance.
func NewVNRecognizedPoint3D() VNRecognizedPoint3D {
	class := getVNRecognizedPoint3DClass()
	rv := objc.Send[VNRecognizedPoint3D](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a point object with the position you specify.
//
// position: The three-dimensional position.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint3D/init(position:)
// position is a [simd.simd_float4x4].
func NewRecognizedPoint3DWithPosition(position objectivec.IObject) VNRecognizedPoint3D {
	instance := getVNRecognizedPoint3DClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPosition:"), position)
	return VNRecognizedPoint3DFromID(rv)
}

// The identifier that provides context about what kind of point the request
// recognizes.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizedPoint3D/identifier
func (r VNRecognizedPoint3D) Identifier() VNRecognizedPointKey {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("identifier"))
	return VNRecognizedPointKey(foundation.NSStringFromID(rv).String())
}

