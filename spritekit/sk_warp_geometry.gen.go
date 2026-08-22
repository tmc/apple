// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKWarpGeometry] class.
var (
	_SKWarpGeometryClass     SKWarpGeometryClass
	_SKWarpGeometryClassOnce sync.Once
)

func getSKWarpGeometryClass() SKWarpGeometryClass {
	_SKWarpGeometryClassOnce.Do(func() {
		_SKWarpGeometryClass = SKWarpGeometryClass{class: objc.GetClass("SKWarpGeometry")}
	})
	return _SKWarpGeometryClass
}

// GetSKWarpGeometryClass returns the class object for SKWarpGeometry.
func GetSKWarpGeometryClass() SKWarpGeometryClass {
	return getSKWarpGeometryClass()
}

type SKWarpGeometryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKWarpGeometryClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKWarpGeometryClass) Alloc() SKWarpGeometry {
	rv := objc.Send[SKWarpGeometry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A definition for a deformation of nodes that conform to [SKWarpable].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometry
type SKWarpGeometry struct {
	objectivec.Object
}

// SKWarpGeometryFromID constructs a [SKWarpGeometry] from an objc.ID.
//
// A definition for a deformation of nodes that conform to [SKWarpable].
func SKWarpGeometryFromID(id objc.ID) SKWarpGeometry {
	return SKWarpGeometry{objectivec.Object{ID: id}}
}

// NOTE: SKWarpGeometry adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKWarpGeometry] class.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometry
type ISKWarpGeometry interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (w SKWarpGeometry) Init() SKWarpGeometry {
	rv := objc.Send[SKWarpGeometry](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w SKWarpGeometry) Autorelease() SKWarpGeometry {
	rv := objc.Send[SKWarpGeometry](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKWarpGeometry creates a new SKWarpGeometry instance.
func NewSKWarpGeometry() SKWarpGeometry {
	class := getSKWarpGeometryClass()
	rv := objc.Send[SKWarpGeometry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w SKWarpGeometry) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](w.ID, objc.Sel("encodeWithCoder:"), coder)
}
