// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSZeroingWeakContainer] class.
var (
	_SLSZeroingWeakContainerClass     SLSZeroingWeakContainerClass
	_SLSZeroingWeakContainerClassOnce sync.Once
)

func getSLSZeroingWeakContainerClass() SLSZeroingWeakContainerClass {
	_SLSZeroingWeakContainerClassOnce.Do(func() {
		_SLSZeroingWeakContainerClass = SLSZeroingWeakContainerClass{class: objc.GetClass("SLSZeroingWeakContainer")}
	})
	return _SLSZeroingWeakContainerClass
}

// GetSLSZeroingWeakContainerClass returns the class object for SLSZeroingWeakContainer.
func GetSLSZeroingWeakContainerClass() SLSZeroingWeakContainerClass {
	return getSLSZeroingWeakContainerClass()
}

type SLSZeroingWeakContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSZeroingWeakContainerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSZeroingWeakContainerClass) Alloc() SLSZeroingWeakContainer {
	rv := objc.Send[SLSZeroingWeakContainer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSZeroingWeakContainer.Invalidate]
//   - [SLSZeroingWeakContainer.GetObject]
//   - [SLSZeroingWeakContainer.Reference]
//   - [SLSZeroingWeakContainer.SetReference]
//   - [SLSZeroingWeakContainer.InitWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSZeroingWeakContainer
type SLSZeroingWeakContainer struct {
	objectivec.Object
}

// SLSZeroingWeakContainerFromID constructs a [SLSZeroingWeakContainer] from an objc.ID.
func SLSZeroingWeakContainerFromID(id objc.ID) SLSZeroingWeakContainer {
	return SLSZeroingWeakContainer{objectivec.Object{ID: id}}
}

// Ensure SLSZeroingWeakContainer implements ISLSZeroingWeakContainer.
var _ ISLSZeroingWeakContainer = SLSZeroingWeakContainer{}

// An interface definition for the [SLSZeroingWeakContainer] class.
//
// # Methods
//
//   - [ISLSZeroingWeakContainer.Invalidate]
//   - [ISLSZeroingWeakContainer.GetObject]
//   - [ISLSZeroingWeakContainer.Reference]
//   - [ISLSZeroingWeakContainer.SetReference]
//   - [ISLSZeroingWeakContainer.InitWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSZeroingWeakContainer
type ISLSZeroingWeakContainer interface {
	objectivec.IObject

	// Topic: Methods

	Invalidate()
	GetObject() objectivec.IObject
	Reference() objectivec.IObject
	SetReference(value objectivec.IObject)
	InitWithObject(object objectivec.IObject) SLSZeroingWeakContainer
}

// Init initializes the instance.
func (s SLSZeroingWeakContainer) Init() SLSZeroingWeakContainer {
	rv := objc.Send[SLSZeroingWeakContainer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSZeroingWeakContainer) Autorelease() SLSZeroingWeakContainer {
	rv := objc.Send[SLSZeroingWeakContainer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSZeroingWeakContainer creates a new SLSZeroingWeakContainer instance.
func NewSLSZeroingWeakContainer() SLSZeroingWeakContainer {
	class := getSLSZeroingWeakContainerClass()
	rv := objc.Send[SLSZeroingWeakContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSZeroingWeakContainer/initWithObject:
func NewSLSZeroingWeakContainerWithObject(object objectivec.IObject) SLSZeroingWeakContainer {
	instance := getSLSZeroingWeakContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:"), object)
	return SLSZeroingWeakContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSZeroingWeakContainer/invalidate
func (s SLSZeroingWeakContainer) Invalidate() {
	objc.Send[objc.ID](s.ID, objc.Sel("invalidate"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSZeroingWeakContainer/initWithObject:
func (s SLSZeroingWeakContainer) InitWithObject(object objectivec.IObject) SLSZeroingWeakContainer {
	rv := objc.Send[SLSZeroingWeakContainer](s.ID, objc.Sel("initWithObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSZeroingWeakContainer/containerWithObject:
func (_SLSZeroingWeakContainerClass SLSZeroingWeakContainerClass) ContainerWithObject(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSZeroingWeakContainerClass.class), objc.Sel("containerWithObject:"), object)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSZeroingWeakContainer/object
func (s SLSZeroingWeakContainer) GetObject() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("object"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSZeroingWeakContainer/reference
func (s SLSZeroingWeakContainer) Reference() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("reference"))
	return objectivec.Object{ID: rv}
}
func (s SLSZeroingWeakContainer) SetReference(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setReference:"), value)
}
