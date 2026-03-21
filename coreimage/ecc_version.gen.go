// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EccVersion] class.
var (
	_EccVersionClass     EccVersionClass
	_EccVersionClassOnce sync.Once
)

func getEccVersionClass() EccVersionClass {
	_EccVersionClassOnce.Do(func() {
		_EccVersionClass = EccVersionClass{class: objc.GetClass("eccVersion")}
	})
	return _EccVersionClass
}

// GetEccVersionClass returns the class object for eccVersion.
func GetEccVersionClass() EccVersionClass {
	return getEccVersionClass()
}

type EccVersionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EccVersionClass) Alloc() EccVersion {
	rv := objc.Send[EccVersion](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/eccVersion-c.ivar
type EccVersion struct {
	objectivec.Object
}

// EccVersionFromID constructs a [EccVersion] from an objc.ID.
func EccVersionFromID(id objc.ID) EccVersion {
	return EccVersion{objectivec.Object{ID: id}}
}
// Ensure EccVersion implements IEccVersion.
var _ IEccVersion = EccVersion{}

// An interface definition for the [EccVersion] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/eccVersion-c.ivar
type IEccVersion interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e EccVersion) Init() EccVersion {
	rv := objc.Send[EccVersion](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EccVersion) Autorelease() EccVersion {
	rv := objc.Send[EccVersion](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEccVersion creates a new EccVersion instance.
func NewEccVersion() EccVersion {
	class := getEccVersionClass()
	rv := objc.Send[EccVersion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

