// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKRegion] class.
var (
	_SKRegionClass     SKRegionClass
	_SKRegionClassOnce sync.Once
)

func getSKRegionClass() SKRegionClass {
	_SKRegionClassOnce.Do(func() {
		_SKRegionClass = SKRegionClass{class: objc.GetClass("SKRegion")}
	})
	return _SKRegionClass
}

// GetSKRegionClass returns the class object for SKRegion.
func GetSKRegionClass() SKRegionClass {
	return getSKRegionClass()
}

type SKRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKRegionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKRegionClass) Alloc() SKRegion {
	rv := objc.Send[SKRegion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// The definition of an arbitrary area.
//
// # Overview
//
// An [SKRegion] object defines a mathematical shape and is typically used to
// determine whether a particular point lies inside this area. For example,
// regions are used to define the area that a physics field can affect.
// Regions are defined using paths and mathematical shapes and can also be
// combined using constructive solid geometry.
//
// # Creating and Initializing Region Objects
//
//   - [SKRegion.InitWithSize]: Initializes a new region with a rectangular area.
//   - [SKRegion.InitWithRadius]: Initializes a new region with a circular area.
//   - [SKRegion.InitWithPath]: Initializes a new region using a Core Graphics path.
//   - [SKRegion.InverseRegion]: Returns a new region that is the mathematical inverse of an existing region.
//   - [SKRegion.RegionByDifferenceFromRegion]: Returns a new region created by subtracting the contents of another region from this region.
//   - [SKRegion.RegionByIntersectionWithRegion]: Returns a new region created by intersecting the contents of this region with another region.
//   - [SKRegion.RegionByUnionWithRegion]: Returns a new region created by combining the contents of this region with another region.
//
// # Interacting with a Region
//
//   - [SKRegion.Path]: Returns a Core Graphics path that defines the region.
//   - [SKRegion.ContainsPoint]: Returns a Boolean value that indicates whether a particular point is contained in the region.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion
type SKRegion struct {
	objectivec.Object
}

// SKRegionFromID constructs a [SKRegion] from an objc.ID.
//
// The definition of an arbitrary area.
func SKRegionFromID(id objc.ID) SKRegion {
	return SKRegion{objectivec.Object{ID: id}}
}

// NOTE: SKRegion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKRegion] class.
//
// # Creating and Initializing Region Objects
//
//   - [ISKRegion.InitWithSize]: Initializes a new region with a rectangular area.
//   - [ISKRegion.InitWithRadius]: Initializes a new region with a circular area.
//   - [ISKRegion.InitWithPath]: Initializes a new region using a Core Graphics path.
//   - [ISKRegion.InverseRegion]: Returns a new region that is the mathematical inverse of an existing region.
//   - [ISKRegion.RegionByDifferenceFromRegion]: Returns a new region created by subtracting the contents of another region from this region.
//   - [ISKRegion.RegionByIntersectionWithRegion]: Returns a new region created by intersecting the contents of this region with another region.
//   - [ISKRegion.RegionByUnionWithRegion]: Returns a new region created by combining the contents of this region with another region.
//
// # Interacting with a Region
//
//   - [ISKRegion.Path]: Returns a Core Graphics path that defines the region.
//   - [ISKRegion.ContainsPoint]: Returns a Boolean value that indicates whether a particular point is contained in the region.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion
type ISKRegion interface {
	objectivec.IObject

	// Topic: Creating and Initializing Region Objects

	// Initializes a new region with a rectangular area.
	InitWithSize(size corefoundation.CGSize) SKRegion
	// Initializes a new region with a circular area.
	InitWithRadius(radius float32) SKRegion
	// Initializes a new region using a Core Graphics path.
	InitWithPath(path coregraphics.CGPathRef) SKRegion
	// Returns a new region that is the mathematical inverse of an existing region.
	InverseRegion() ISKRegion
	// Returns a new region created by subtracting the contents of another region from this region.
	RegionByDifferenceFromRegion(region ISKRegion) ISKRegion
	// Returns a new region created by intersecting the contents of this region with another region.
	RegionByIntersectionWithRegion(region ISKRegion) ISKRegion
	// Returns a new region created by combining the contents of this region with another region.
	RegionByUnionWithRegion(region ISKRegion) ISKRegion

	// Topic: Interacting with a Region

	// Returns a Core Graphics path that defines the region.
	Path() coregraphics.CGPathRef
	// Returns a Boolean value that indicates whether a particular point is contained in the region.
	ContainsPoint(point corefoundation.CGPoint) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r SKRegion) Init() SKRegion {
	rv := objc.Send[SKRegion](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r SKRegion) Autorelease() SKRegion {
	rv := objc.Send[SKRegion](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKRegion creates a new SKRegion instance.
func NewSKRegion() SKRegion {
	class := getSKRegionClass()
	rv := objc.Send[SKRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new region using a Core Graphics path.
//
// path: A path that defines the new region’s shape. The path is assumed to use
// the even-odd winding rule.
//
// # Return Value
//
// A newly initialized region.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(path:)
func NewRegionWithPath(path coregraphics.CGPathRef) SKRegion {
	instance := getSKRegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPath:"), path)
	return SKRegionFromID(rv)
}

// Initializes a new region with a circular area.
//
// radius: The radius of the region in points.
//
// # Return Value
//
// A newly initialized region. The region is circular and centered on the
// origin.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(radius:)
func NewRegionWithRadius(radius float32) SKRegion {
	instance := getSKRegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRadius:"), radius)
	return SKRegionFromID(rv)
}

// Initializes a new region with a rectangular area.
//
// size: The size of the rectangle in points.
//
// # Return Value
//
// A newly initialized region. The region is rectangular and centered on the
// origin.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(size:)
func NewRegionWithSize(size corefoundation.CGSize) SKRegion {
	instance := getSKRegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:"), size)
	return SKRegionFromID(rv)
}

// Initializes a new region with a rectangular area.
//
// size: The size of the rectangle in points.
//
// # Return Value
//
// A newly initialized region. The region is rectangular and centered on the
// origin.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(size:)
func (r SKRegion) InitWithSize(size corefoundation.CGSize) SKRegion {
	rv := objc.Send[SKRegion](r.ID, objc.Sel("initWithSize:"), size)
	return rv
}

// Initializes a new region with a circular area.
//
// radius: The radius of the region in points.
//
// # Return Value
//
// A newly initialized region. The region is circular and centered on the
// origin.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(radius:)
func (r SKRegion) InitWithRadius(radius float32) SKRegion {
	rv := objc.Send[SKRegion](r.ID, objc.Sel("initWithRadius:"), radius)
	return rv
}

// Initializes a new region using a Core Graphics path.
//
// path: A path that defines the new region’s shape. The path is assumed to use
// the even-odd winding rule.
//
// # Return Value
//
// A newly initialized region.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(path:)
func (r SKRegion) InitWithPath(path coregraphics.CGPathRef) SKRegion {
	rv := objc.Send[SKRegion](r.ID, objc.Sel("initWithPath:"), path)
	return rv
}

// Returns a new region that is the mathematical inverse of an existing
// region.
//
// # Return Value
//
// A new region object whose contents include all points that are not in the
// current region.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/inverse()
func (r SKRegion) InverseRegion() ISKRegion {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("inverseRegion"))
	return SKRegionFromID(rv)
}

// Returns a new region created by subtracting the contents of another region
// from this region.
//
// region: The region to subtract.
//
// # Return Value
//
// A new region whose contents include all points in the current region that
// are not also included in the second region.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/byDifference(from:)
func (r SKRegion) RegionByDifferenceFromRegion(region ISKRegion) ISKRegion {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("regionByDifferenceFromRegion:"), region)
	return SKRegionFromID(rv)
}

// Returns a new region created by intersecting the contents of this region
// with another region.
//
// region: The region to intersect.
//
// # Return Value
//
// A new region whose contents include all points that are included in both
// regions.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/byIntersection(with:)
func (r SKRegion) RegionByIntersectionWithRegion(region ISKRegion) ISKRegion {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("regionByIntersectionWithRegion:"), region)
	return SKRegionFromID(rv)
}

// Returns a new region created by combining the contents of this region with
// another region.
//
// region: The region to combine with the current region.
//
// # Return Value
//
// A new region whose contents include all points that are included in either
// region.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/byUnion(with:)
func (r SKRegion) RegionByUnionWithRegion(region ISKRegion) ISKRegion {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("regionByUnionWithRegion:"), region)
	return SKRegionFromID(rv)
}

// Returns a Boolean value that indicates whether a particular point is
// contained in the region.
//
// point: A point.
//
// # Return Value
//
// true if the point is contained in the region; otherwise, false.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/contains(_:)
func (r SKRegion) ContainsPoint(point corefoundation.CGPoint) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("containsPoint:"), point)
	return rv
}
func (r SKRegion) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a region that defines a region that includes all points.
//
// # Return Value
//
// Returns a singleton region that covers an infinite area.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/infinite()
func (_SKRegionClass SKRegionClass) InfiniteRegion() SKRegion {
	rv := objc.Send[objc.ID](objc.ID(_SKRegionClass.class), objc.Sel("infiniteRegion"))
	return SKRegionFromID(rv)
}

// Returns a Core Graphics path that defines the region.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRegion/path
func (r SKRegion) Path() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](r.ID, objc.Sel("path"))
	return coregraphics.CGPathRef(rv)
}
