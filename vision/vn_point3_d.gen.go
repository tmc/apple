// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNPoint3D] class.
var (
	_VNPoint3DClass     VNPoint3DClass
	_VNPoint3DClassOnce sync.Once
)

func getVNPoint3DClass() VNPoint3DClass {
	_VNPoint3DClassOnce.Do(func() {
		_VNPoint3DClass = VNPoint3DClass{class: objc.GetClass("VNPoint3D")}
	})
	return _VNPoint3DClass
}

// GetVNPoint3DClass returns the class object for VNPoint3D.
func GetVNPoint3DClass() VNPoint3DClass {
	return getVNPoint3DClass()
}

type VNPoint3DClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNPoint3DClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNPoint3DClass) Alloc() VNPoint3D {
	rv := objc.Send[VNPoint3D](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a 3D point in an image.
//
// # Creating a Point
//
//   - [VNPoint3D.InitWithPosition]: Creates a point object with the position you specify.
//
// # Getting the Position
//
//   - [VNPoint3D.Position]: The three-dimensional position.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint3D
type VNPoint3D struct {
	objectivec.Object
}

// VNPoint3DFromID constructs a [VNPoint3D] from an objc.ID.
//
// An object that represents a 3D point in an image.
func VNPoint3DFromID(id objc.ID) VNPoint3D {
	return VNPoint3D{objectivec.Object{ID: id}}
}
// NOTE: VNPoint3D adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNPoint3D] class.
//
// # Creating a Point
//
//   - [IVNPoint3D.InitWithPosition]: Creates a point object with the position you specify.
//
// # Getting the Position
//
//   - [IVNPoint3D.Position]: The three-dimensional position.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint3D
type IVNPoint3D interface {
	objectivec.IObject

	// Topic: Creating a Point

	// Creates a point object with the position you specify.
	InitWithPosition(position objectivec.IObject) VNPoint3D

	// Topic: Getting the Position

	// The three-dimensional position.
	Position() objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p VNPoint3D) Init() VNPoint3D {
	rv := objc.Send[VNPoint3D](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p VNPoint3D) Autorelease() VNPoint3D {
	rv := objc.Send[VNPoint3D](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNPoint3D creates a new VNPoint3D instance.
func NewVNPoint3D() VNPoint3D {
	class := getVNPoint3DClass()
	rv := objc.Send[VNPoint3D](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a point object with the position you specify.
//
// position: The three-dimensional position.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint3D/init(position:)
// position is a [simd.simd_float4x4].
func NewPoint3DWithPosition(position objectivec.IObject) VNPoint3D {
	instance := getVNPoint3DClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPosition:"), position)
	return VNPoint3DFromID(rv)
}

// Creates a point object with the position you specify.
//
// position: The three-dimensional position.
//
// position is a [simd.simd_float4x4].
//
// See: https://developer.apple.com/documentation/Vision/VNPoint3D/init(position:)
// position is a [simd.simd_float4x4].
func (p VNPoint3D) InitWithPosition(position objectivec.IObject) VNPoint3D {
	rv := objc.Send[VNPoint3D](p.ID, objc.Sel("initWithPosition:"), position)
	return rv
}
func (p VNPoint3D) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The three-dimensional position.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint3D/position
func (p VNPoint3D) Position() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("position"))
	return objectivec.Object{ID: rv}
}

