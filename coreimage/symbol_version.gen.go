// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SymbolVersion] class.
var (
	_SymbolVersionClass     SymbolVersionClass
	_SymbolVersionClassOnce sync.Once
)

func getSymbolVersionClass() SymbolVersionClass {
	_SymbolVersionClassOnce.Do(func() {
		_SymbolVersionClass = SymbolVersionClass{class: objc.GetClass("symbolVersion")}
	})
	return _SymbolVersionClass
}

// GetSymbolVersionClass returns the class object for symbolVersion.
func GetSymbolVersionClass() SymbolVersionClass {
	return getSymbolVersionClass()
}

type SymbolVersionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SymbolVersionClass) Alloc() SymbolVersion {
	rv := objc.Send[SymbolVersion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/symbolVersion-c.ivar
type SymbolVersion struct {
	objectivec.Object
}

// SymbolVersionFromID constructs a [SymbolVersion] from an objc.ID.
func SymbolVersionFromID(id objc.ID) SymbolVersion {
	return SymbolVersion{objectivec.Object{ID: id}}
}
// Ensure SymbolVersion implements ISymbolVersion.
var _ ISymbolVersion = SymbolVersion{}

// An interface definition for the [SymbolVersion] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/symbolVersion-c.ivar
type ISymbolVersion interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SymbolVersion) Init() SymbolVersion {
	rv := objc.Send[SymbolVersion](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SymbolVersion) Autorelease() SymbolVersion {
	rv := objc.Send[SymbolVersion](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSymbolVersion creates a new SymbolVersion instance.
func NewSymbolVersion() SymbolVersion {
	class := getSymbolVersionClass()
	rv := objc.Send[SymbolVersion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

