// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMemory] class.
var (
	_VZMemoryClass     VZMemoryClass
	_VZMemoryClassOnce sync.Once
)

func getVZMemoryClass() VZMemoryClass {
	_VZMemoryClassOnce.Do(func() {
		_VZMemoryClass = VZMemoryClass{class: objc.GetClass("_VZMemory")}
	})
	return _VZMemoryClass
}

// GetVZMemoryClass returns the class object for _VZMemory.
func GetVZMemoryClass() VZMemoryClass {
	return getVZMemoryClass()
}

type VZMemoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMemoryClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMemoryClass) Alloc() VZMemory {
	rv := objc.SendIfResponds[VZMemory](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZMemory struct {
	VZGuestMemoryMapping
}

// VZMemoryFromID constructs a [VZMemory] from an objc.ID.
func VZMemoryFromID(id objc.ID) VZMemory {
	return VZMemory{VZGuestMemoryMapping: VZGuestMemoryMappingFromID(id)}
}

// Ensure VZMemory implements IVZMemory.
var _ IVZMemory = VZMemory{}

// An interface definition for the [VZMemory] class.
type IVZMemory interface {
	IVZGuestMemoryMapping
}

// Init initializes the instance.
func (v VZMemory) Init() VZMemory {
	rv := objc.SendIfResponds[VZMemory](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMemory) Autorelease() VZMemory {
	rv := objc.SendIfResponds[VZMemory](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemory creates a new VZMemory instance.
func NewVZMemory() VZMemory {
	class := getVZMemoryClass()
	rv := objc.SendIfResponds[VZMemory](objc.ID(class.class), objc.Sel("new"))
	return rv
}
