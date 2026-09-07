// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBootLoader] class.
var (
	_VZBootLoaderClass     VZBootLoaderClass
	_VZBootLoaderClassOnce sync.Once
)

func getVZBootLoaderClass() VZBootLoaderClass {
	_VZBootLoaderClassOnce.Do(func() {
		_VZBootLoaderClass = VZBootLoaderClass{class: objc.GetClass("VZBootLoader")}
	})
	return _VZBootLoaderClass
}

// GetVZBootLoaderClass returns the class object for VZBootLoader.
func GetVZBootLoaderClass() VZBootLoaderClass {
	return getVZBootLoaderClass()
}

type VZBootLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBootLoaderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBootLoaderClass) Alloc() VZBootLoader {
	rv := objc.SendIfResponds[VZBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZBootLoader._init]
//   - [VZBootLoader.DebugDescription]
//   - [VZBootLoader.Description]
//   - [VZBootLoader.Hash]
//   - [VZBootLoader.Superclass]
type VZBootLoader struct {
	objectivec.Object
}

// VZBootLoaderFromID constructs a [VZBootLoader] from an objc.ID.
func VZBootLoaderFromID(id objc.ID) VZBootLoader {
	return VZBootLoader{objectivec.Object{ID: id}}
}

// Ensure VZBootLoader implements IVZBootLoader.
var _ IVZBootLoader = VZBootLoader{}

// An interface definition for the [VZBootLoader] class.
//
// # Methods
//
//   - [IVZBootLoader._init]
//   - [IVZBootLoader.DebugDescription]
//   - [IVZBootLoader.Description]
//   - [IVZBootLoader.Hash]
//   - [IVZBootLoader.Superclass]
type IVZBootLoader interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZBootLoader) Init() VZBootLoader {
	rv := objc.SendIfResponds[VZBootLoader](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBootLoader) Autorelease() VZBootLoader {
	rv := objc.SendIfResponds[VZBootLoader](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBootLoader creates a new VZBootLoader instance.
func NewVZBootLoader() VZBootLoader {
	class := getVZBootLoaderClass()
	rv := objc.SendIfResponds[VZBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZBootLoader) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZBootLoader) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZBootLoader) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZBootLoader) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZBootLoader) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
