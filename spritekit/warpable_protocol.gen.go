// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for objects that can be warped and animated by an [SKWarpGeometry](<https://developer.apple.com/documentation/SpriteKit/SKWarpGeometry>).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable
type SKWarpable interface {
	objectivec.IObject

	// The maximum number of subdivision iterations used to generate the final vertices.
	//
	// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/subdivisionLevels
	SubdivisionLevels() int
	SetSubdivisionLevels(value int)

	// The warp geometry used to define the distortion.
	//
	// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/warpGeometry
	WarpGeometry() ISKWarpGeometry
	SetWarpGeometry(value ISKWarpGeometry)
}

// SKWarpableObject wraps an existing Objective-C object that conforms to the SKWarpable protocol.
type SKWarpableObject struct {
	objectivec.Object
}

func (o SKWarpableObject) BaseObject() objectivec.Object {
	return o.Object
}

// SKWarpableObjectFromID constructs a [SKWarpableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SKWarpableObjectFromID(id objc.ID) SKWarpableObject {
	return SKWarpableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The maximum number of subdivision iterations used to generate the final
// vertices.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/subdivisionLevels
func (o SKWarpableObject) SubdivisionLevels() int {
	rv := objc.Send[int](o.ID, objc.Sel("subdivisionLevels"))
	return int(rv)
}

func (o SKWarpableObject) SetSubdivisionLevels(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setSubdivisionLevels:"), value)
}

// The warp geometry used to define the distortion.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/warpGeometry
func (o SKWarpableObject) WarpGeometry() ISKWarpGeometry {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("warpGeometry"))
	return SKWarpGeometryFromID(rv)
}

func (o SKWarpableObject) SetWarpGeometry(value ISKWarpGeometry) {
	objc.Send[struct{}](o.ID, objc.Sel("setWarpGeometry:"), value)
}
